package io.greenerthumb.collection;

import org.junit.Test;

import static org.junit.Assert.*;

public class ArrayViewUnitTest {
    @Test
    public void sizeOfViewIsCorrect() {
        ArrayView<Integer> view = new ArrayView<>(new Integer[]{0, 0, 0});
        assertEquals(view.size(), 3);
    }

    @Test
    public void valueAtIndexIsCorrect() {
        ArrayView<Integer> view = new ArrayView<>(new Integer[]{0, 1, 2});
        assertEquals(new Integer(0), view.at(0));
        assertEquals(new Integer(1), view.at(1));
        assertEquals(new Integer(2), view.at(2));
    }

    @Test
    public void arrayIsCorrect() {
        ArrayView<Integer> view = new ArrayView<>(new Integer[]{0, 1, 2});
        Integer[] expected = new Integer[]{0, 1, 2};
        Integer[] actual = view.array();
        assertArrayEquals(expected, actual);
    }

    @Test
    public void arrayIsCopy() {
        Integer[] original = new Integer[]{0};
        ArrayView<Integer> view = new ArrayView<>(original);
        Integer[] copy = view.array();
        copy[0]++;
        assertEquals(new Integer(0), original[0]);
        assertEquals(new Integer(1), copy[0]);
    }

    @Test
    public void viewOf() {
        ArrayView<Integer> view = new ArrayView<>(new Integer[]{1, 2, 3, 4});
        ArrayView<Integer> leftHalf = view.viewOf(0, 2);
        ArrayView<Integer> rightHalf = view.viewOf(2, 2);
        assertEquals(2, leftHalf.size());
        assertEquals(2, rightHalf.size());
        assertEquals(new Integer(1), leftHalf.at(0));
        assertEquals(new Integer(2), leftHalf.at(1));
        assertEquals(new Integer(3), rightHalf.at(0));
        assertEquals(new Integer(4), rightHalf.at(1));
    }

    @Test
    public void viewOfView() {
        ArrayView<Integer> view = new ArrayView<>(new Integer[]{1, 2, 3, 4});
        view = view.viewOf(2, 2);
        view = view.viewOf(1, 1);
        assertEquals(1, view.size());
        assertEquals(new Integer(4), view.at(0));
    }

    @Test
    public void atIndexOutOfBounds() {
        ArrayView<Integer> view = new ArrayView<>(new Integer[]{});
        Exception exception = null;
        try {
            view.at(0);
        } catch (Exception actual) {
            exception = actual;
        }
        assertNotNull(exception);
        assertEquals(ArrayIndexOutOfBoundsException.class, exception.getClass());
    }

    @Test
    public void arrayIndexOutOfBounds() {
        ArrayView<Integer> view = new ArrayView<>(new Integer[]{});
        Exception exception = null;
        try {
            view.viewOf(0, 1).array();
        } catch (Exception actual) {
            exception = actual;
        }
        assertNotNull(exception);
        assertEquals(ArrayIndexOutOfBoundsException.class, exception.getClass());
    }

    @Test
    public void advance() {
        ArrayView<Integer> view = new ArrayView<>(new Integer[]{1, 2, 3});
        assertEquals(new Integer(1), view.at(0));
        view = ArrayView.advance(view, 1);
        assertEquals(new Integer(2), view.at(0));
        view = ArrayView.advance(view, 1);
        assertEquals(new Integer(3), view.at(0));
    }
}
