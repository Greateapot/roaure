package re.greateapot.roaure.models;

import io.grpc.Status;

public class StatusWithCallback {

    public interface Callback {
        void retry();
    }

    public final Status status;
    public final Callback callback;

    public StatusWithCallback(Status status, Callback callback) {
        this.status = status;
        this.callback = callback;
    }
}
