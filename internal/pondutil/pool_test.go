package pondutil

import (
	"context"
	"fmt"
	"sort"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewPoolStop(t *testing.T) {
	p := New(2, 3)
	assert.Equal(t, 2, p.MaxWorkers())
	assert.Equal(t, 3, p.MaxCapacity())
	p.Stop()
	time.Sleep(time.Millisecond)
	assert.Equal(t, true, p.Stopped())
}

func TestRunTasks(t *testing.T) {
	p := New(2, 3)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	t.Run("run tasks", func(t *testing.T) {
		var tasks []Task[int64]
		tasks = append(tasks,
			func() (int64, error) {
				time.Sleep(20 * time.Millisecond)
				return int64(1), nil
			},
			func() (int64, error) {
				time.Sleep(30 * time.Millisecond)
				return int64(2), nil
			},
			func() (int64, error) {
				time.Sleep(10 * time.Millisecond)
				return int64(3), nil
			},
		)

		result, err := RunTasks[int64](p, ctx, tasks...)
		assert.NoError(t, err)

		// sorting just for assert
		sort.Slice(result, func(i, j int) bool {
			return result[i] < result[j]
		})
		assert.Equal(t, len(tasks), len(result))
		assert.Equal(t, []int64{1, 2, 3}, result)
	})

	t.Run("run tasks with error", func(t *testing.T) {
		var tasks []Task[int64]
		tasks = append(tasks,
			func() (int64, error) {
				time.Sleep(20 * time.Millisecond)
				return int64(1), nil
			},
			func() (int64, error) {
				time.Sleep(30 * time.Millisecond)
				return 0, fmt.Errorf("custom error")
			},
			func() (int64, error) {
				time.Sleep(10 * time.Millisecond)
				return int64(3), nil
			},
		)

		result, err := RunTasks[int64](p, ctx, tasks...)

		assert.Error(t, err)
		assert.Equal(t, err.Error(), "custom error")

		assert.Equal(t, 0, len(result))
	})

	t.Run("run tasks with custom type struct", func(t *testing.T) {
		type ResultType struct {
			a int64
		}
		var tasks []Task[ResultType]
		tasks = append(tasks,
			func() (ResultType, error) {
				time.Sleep(20 * time.Millisecond)
				return ResultType{a: 1}, nil
			},
			func() (ResultType, error) {
				time.Sleep(30 * time.Millisecond)
				return ResultType{a: 2}, nil
			},
			func() (ResultType, error) {
				time.Sleep(10 * time.Millisecond)
				return ResultType{a: 3}, nil
			},
		)

		result, err := RunTasks(p, ctx, tasks...)

		assert.NoError(t, err)

		assert.Equal(t, len(tasks), len(result))
		expected := []ResultType{
			{a: 1},
			{a: 2},
			{a: 3},
		}
		sort.Slice(result, func(i, j int) bool {
			return result[i].a < result[j].a
		})
		assert.Equal(t, expected, result)
	})

	t.Run("run tasks with channel", func(t *testing.T) {
		ch := make(chan Task[int64])
		wg := sync.WaitGroup{}
		wg.Add(1)
		go func() {
			result, err := RunTasksWithSupplierChan[int64](p, ctx, ch)
			assert.NoError(t, err)

			// sorting just for assert
			sort.Slice(result, func(i, j int) bool {
				return result[i] < result[j]
			})
			assert.Equal(t, 3, len(result))
			assert.Equal(t, []int64{1, 2, 3}, result)
			wg.Done()
		}()
		ch <- func() (int64, error) {
			time.Sleep(20 * time.Millisecond)
			return int64(1), nil
		}
		ch <- func() (int64, error) {
			time.Sleep(30 * time.Millisecond)
			return int64(2), nil
		}
		ch <- func() (int64, error) {
			time.Sleep(10 * time.Millisecond)
			return int64(3), nil
		}
		close(ch)
		wg.Wait()
	})

	t.Run("run tasks with provider func", func(t *testing.T) {
		result, err := RunTasksWithSupplierFunc(p, ctx, func(out chan<- Task[int64]) {
			for i := 1; i <= 3; i++ {
				v := i
				out <- func() (int64, error) {
					return int64(v), nil
				}
			}
		})
		assert.NoError(t, err)

		// sorting just for assert
		sort.Slice(result, func(i, j int) bool {
			return result[i] < result[j]
		})
		assert.Equal(t, 3, len(result))
		assert.Equal(t, []int64{1, 2, 3}, result)
	})

	t.Run("run task group with supplier", func(t *testing.T) {
		tg, _ := p.GroupContext(ctx)
		result, err := RunTaskGroupWithSupplierFunc(tg, func(out chan<- Task[int64]) {
			for i := 1; i <= 3; i++ {
				v := i
				out <- func() (int64, error) {
					return int64(v), nil
				}
			}
		})
		assert.NoError(t, err)

		// sorting just for assert
		sort.Slice(result, func(i, j int) bool {
			return result[i] < result[j]
		})
		assert.Equal(t, 3, len(result))
		assert.Equal(t, []int64{1, 2, 3}, result)
	})
}
