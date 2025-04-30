import time
import multiprocessing
import sys

TOTAL_ITERATIONS = 100_000_000
NUM_PROCESSES = 8


def compute_range(start, end):
    result = 0
    for i in range(start, end):
        val = (i * 31) ^ (i + 17)
        result ^= val
    return result


def single_threaded():
    start = time.perf_counter()
    result = compute_range(0, TOTAL_ITERATIONS)
    end = time.perf_counter()
    print(f"Single-threaded: {end - start:.4f} seconds\n")
    print(f"Result checksum: {result}")


def worker(start, end, return_dict, idx):
    return_dict[idx] = compute_range(start, end)


def multi_process():
    manager = multiprocessing.Manager()
    return_dict = manager.dict()
    processes = []
    chunk_size = TOTAL_ITERATIONS // NUM_PROCESSES

    start = time.perf_counter()

    for i in range(NUM_PROCESSES):
        chunk_start = i * chunk_size
        chunk_end = chunk_start + chunk_size
        p = multiprocessing.Process(
            target=worker, args=(chunk_start, chunk_end, return_dict, i)
        )
        processes.append(p)
        p.start()

    for p in processes:
        p.join()

    result = 0
    for i in range(NUM_PROCESSES):
        result ^= return_dict[i]

    end = time.perf_counter()
    print(f"Multi-process: {end - start:.4f} seconds\n")
    print(f"Result checksum: {result}")


if __name__ == "__main__":
    print(sys.version)
    print("Running compute benchmark with", TOTAL_ITERATIONS, "iterations...\n")
    single_threaded()
    multi_process()

