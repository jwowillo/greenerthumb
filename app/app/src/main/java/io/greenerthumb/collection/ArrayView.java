package io.greenerthumb.collection;

import java.util.Arrays;

/**
 * ArrayView is an immutable view of an array.
 *
 * @param <T> Type of the array to view.
 */
public class ArrayView<T> {
    private final T[] array;
    private final int start;
    private final int size;

    /**
     * @param array to view.
     */
    public ArrayView(T[] array) {
        this.array = array;
        this.start = 0;
        this.size = array.length;
    }

    private ArrayView(T[] array, int start, int size) {
        this.array = array;
        this.start = start;
        this.size = size;
    }

    /**
     * @param i is the index.
     * @return The value at the index.
     */
    public T at(int i) {
        return array[start+i];
    }

    /**
     * @return size of the view.
     */
    public int size() {
        return size;
    }

    /**
     * @param start of the range.
     * @param size of the range.
     * @return viewOf the passed range.
     */
    public ArrayView<T> viewOf(int start, int size) {
        return new ArrayView<>(array, this.start + start, size);
    }

    /**
     * @return array induced by the view.
     */
    public T[] array() {
        if (array.length < start + size) {
            throw new ArrayIndexOutOfBoundsException();
        }
        return Arrays.copyOfRange(array, start, start + size);
    }

    /**
     * @param view to advance.
     * @param count to advance.
     * @param <T> Type of view.
     * @return The advanced view.
     */
    public static <T> ArrayView<T> advance(ArrayView<T> view, int count) {
        return view.viewOf(count, view.size() - count);
    }
}
