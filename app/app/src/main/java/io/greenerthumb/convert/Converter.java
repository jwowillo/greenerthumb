package io.greenerthumb.convert;

import java.util.Optional;

/**
 * Converter converts a T to a U if it is possible.
 *
 * An empty Optional is returned otherwise.
 *
 * @param <T> Type to convert from.
 * @param <U> Type to convert to.
 */
@FunctionalInterface
public interface Converter<T, U> {
    Optional<U> convert(T t);
}
