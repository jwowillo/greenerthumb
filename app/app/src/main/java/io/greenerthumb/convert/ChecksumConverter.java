package io.greenerthumb.convert;

import java.util.Optional;

import io.greenerthumb.collection.ArrayView;

/**
 * ChecksumConverter tries to compare the checksum in the first 4 bytes of an
 * ArrayView<Byte> to the actual sum of the rest and retursn the rest if the
 * sums match.
 *
 * An empty Optional is returned if the sums don't match.
 */
public class ChecksumConverter implements Converter<ArrayView<Byte>, ArrayView<Byte>> {
    @Override
    public Optional<ArrayView<Byte>> convert(ArrayView<Byte> data) {
        if (data.size() < 4) {
            return Optional.empty(); // Too small for a 4-byte checksum.
        }

        int checksum = intFrom(data);

        ArrayView<Byte> rest = ArrayView.advance(data, 4);

        int sum = sumOf(rest);

        if (checksum != sum) {
            return Optional.empty();
        }

        return Optional.of(rest);
    }

    private static int intFrom(ArrayView<Byte> data) {
        int byte0 = data.at(0);
        int byte1 = data.at(1);
        int byte2 = data.at(2);
        int byte3 = data.at(3);
        if (byte0 < 0) {
            byte0 += 0xff + 1;
        }
        if (byte1 < 0) {
            byte1 += 0xff + 1;
        }
        if (byte2 < 0) {
            byte2 += 0xff + 1;
        }
        if (byte3 < 0) {
            byte3 += 0xff + 1;
        }
        return byte3 |
                byte2 << 8 |
                byte1 << 16 |
                byte0 << 24;
    }

    private static int sumOf(ArrayView<Byte> data) {
        int sum = 0;
        for (int i = 0; i < data.size(); i++) {
            int value = data.at(i);
            if (value < 0) {
                value += 0xff + 1;
            }
            sum += value;
        }
        return sum;
    }
}