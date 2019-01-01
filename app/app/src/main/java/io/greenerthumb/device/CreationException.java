package io.greenerthumb.device;

/**
 * CreationException occurs while creating an object.
 */
public class CreationException extends Exception {
    public CreationException(String reason) {
        super(reason);
    }
}
