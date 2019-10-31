package _46_lru_cache

import "testing"

func Test(t *testing.T) {
	inputHandle := []string{"LRUCache", "put", "put", "put", "put", "put", "get", "put", "get", "get", "put", "get", "put", "put", "put", "get", "put", "get", "get", "get", "get", "put", "put", "get", "get", "get", "put", "put", "get", "put", "get", "put", "get", "get", "get", "put", "put", "put", "get", "put", "get", "get", "put", "put", "get", "put", "put", "put", "put", "get", "put", "put", "get", "put", "put", "get", "put", "put", "put", "put", "put", "get", "put", "put", "get", "put", "get", "get", "get", "put", "get", "get", "put", "put", "put", "put", "get", "put", "put", "put", "put", "get", "get", "get", "put", "put", "put", "get", "put", "put", "put", "get", "put", "put", "put", "get", "get", "get", "put", "put", "put", "put", "get", "put", "put", "put", "put", "put", "put", "put"}
	inputData := [][]int{{10}, {10, 13}, {3, 17}, {6, 11}, {10, 5}, {9, 10}, {13}, {2, 19}, {2}, {3}, {5, 25}, {8}, {9, 22}, {5, 5}, {1, 30}, {11}, {9, 12}, {7}, {5}, {8}, {9}, {4, 30}, {9, 3}, {9}, {10}, {10}, {6, 14}, {3, 1}, {3}, {10, 11}, {8}, {2, 14}, {1}, {5}, {4}, {11, 4}, {12, 24}, {5, 18}, {13}, {7, 23}, {8}, {12}, {3, 27}, {2, 12}, {5}, {2, 9}, {13, 4}, {8, 18}, {1, 7}, {6}, {9, 29}, {8, 21}, {5}, {6, 30}, {1, 12}, {10}, {4, 15}, {7, 22}, {11, 26}, {8, 17}, {9, 29}, {5}, {3, 4}, {11, 30}, {12}, {4, 29}, {3}, {9}, {6}, {3, 4}, {1}, {10}, {3, 29}, {10, 28}, {1, 20}, {11, 13}, {3}, {3, 12}, {3, 8}, {10, 9}, {3, 26}, {8}, {7}, {5}, {13, 17}, {2, 27}, {11, 15}, {12}, {9, 19}, {2, 15}, {3, 16}, {1}, {12, 17}, {9, 1}, {6, 19}, {4}, {5}, {5}, {8, 1}, {11, 7}, {5, 2}, {9, 28}, {1}, {2, 2}, {7, 4}, {4, 22}, {7, 24}, {9, 26}, {13, 28}, {11, 26}}

	expect := []int{-100, -100, -100, -100, -100, -100, -1, -100, 19, 17, -100, -1, -100, -100, -100, -1, -100, -1, 5, -1, 12, -100, -100, 3, 5, 5, -100, -100, 1, -100, -1, -100, 30, 5, 30, -100, -100, -100, -1, -100, -1, 24, -100, -100, 18, -100, -100, -100, -100, -1, -100, -100, 18, -100, -100, -1, -100, -100, -100, -100, -100, 18, -100, -100, -1, -100, 4, 29, 30, -100, 12, -1, -100, -100, -100, -100, 29, -100, -100, -100, -100, 17, 22, 18, -100, -100, -100, -1, -100, -100, -100, 20, -100, -100, -100, -1, 18, 18, -100, -100, -100, -100, 20, -100, -100, -100, -100, -100, -100, -100}
	test(inputHandle, inputData, expect, t)
}

func test(inputHandle []string, inputData [][]int, expect []int, t *testing.T) {
	length := len(inputData)
	var cache *LRUCache

	for i := 0; i < length; i++ {
		switch inputHandle[i] {
		case "LRUCache":
			capacity := inputData[i][0]
			tmp := Constructor(capacity)
			cache = &tmp
			break
		case "put":
			key := inputData[i][0]
			value := inputData[i][1]
			cache.Put(key, value)
			break
		case "get":
			key := inputData[i][0]
			v := cache.Get(key)
			if v != expect[i] {
				t.Fail()
			}
			break
		}
	}
}
