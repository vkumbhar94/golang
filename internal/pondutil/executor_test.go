package pondutil

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRunTasksAndPipe(t *testing.T) {
	p := New(3, 5)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	t.Run("run tasks and pipe", func(t *testing.T) {
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

		res, err := RunTasksAndPipe(p, ctx, tasks...)
		assert.NoError(t, err)
		var result []int64
		for r := range res {
			result = append(result, r)

		}

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
		defer func() { _ = recover() }()
		res, err := RunTasksAndPipe(p, ctx, tasks...)

		assert.NoError(t, err)
		//assert.Equal(t, err.Error(), "custom error")

		var result []int64
		for r := range res {
			result = append(result, r)
		}
		assert.Equal(t, true, len(result) < 3)
	})
	t.Run("chained tasks", func(t *testing.T) {
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

		firstTasksChan, err := RunTasksAndPipe(p, ctx, tasks...)
		assert.NoError(t, err)

		inch := MapToTasks(firstTasksChan, func(i int64) Task[string] {
			return func() (string, error) {
				if i <= 0 {
					return "", fmt.Errorf("invalid repeat length")
				}
				return strings.Repeat(fmt.Sprint(i), int(i)), nil
			}
		})

		res, err := RunTasksWithSupplierChanAndPipe(p, ctx, inch)
		var result []string
		for r := range res {
			result = append(result, r)
		}
		// sorting just for assert
		sort.Slice(result, func(i, j int) bool {
			return result[i] < result[j]
		})
		assert.Equal(t, len(tasks), len(result))
		assert.Equal(t, []string{"1", "22", "333"}, result)
	})
}
