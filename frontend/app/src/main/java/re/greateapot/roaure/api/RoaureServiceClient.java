package re.greateapot.roaure.api;


import com.google.protobuf.Empty;

import io.grpc.ManagedChannel;
import io.grpc.Status;
import io.grpc.android.AndroidChannelBuilder;
import io.grpc.stub.StreamObserver;

import re.greateapot.roaure.api.dto.CreateScheduleRequest;
import re.greateapot.roaure.api.dto.DeleteScheduleRequest;
import re.greateapot.roaure.api.dto.GetMetricsRequest;
import re.greateapot.roaure.api.dto.IperfServerConf;
import re.greateapot.roaure.api.dto.ListSchedulesResponse;
import re.greateapot.roaure.api.dto.Metrics;
import re.greateapot.roaure.api.dto.MonitorConf;
import re.greateapot.roaure.api.dto.RoaureServiceGrpc;
import re.greateapot.roaure.api.dto.RouterConf;
import re.greateapot.roaure.api.dto.Schedule;
import re.greateapot.roaure.api.dto.UpdateIperfServerConfRequest;
import re.greateapot.roaure.api.dto.UpdateMonitorConfRequest;
import re.greateapot.roaure.api.dto.UpdateRouterConfRequest;
import re.greateapot.roaure.api.dto.UpdateScheduleRequest;

public class RoaureServiceClient {


    public interface OnValueListener<T> {
        void onValue(T value);
    }

    public interface OnStatusListener {
        void onStatus(Status status);
    }

    public interface OnCompletedListener {
        void onCompleted();
    }

    private static volatile RoaureServiceClient instance;
    private final RoaureServiceGrpc.RoaureServiceStub asyncStub;
    private final ManagedChannel channel;

    private RoaureServiceClient(String host, int port) {
        this(AndroidChannelBuilder.forAddress(host, port).usePlaintext().build());
    }

    private RoaureServiceClient(ManagedChannel channel) {
        this.channel = channel;
        this.asyncStub = RoaureServiceGrpc.newStub(channel);
    }

    public static synchronized void init(String host, int port) {
        if (instance == null) {
            instance = new RoaureServiceClient(host, port);
        }
    }

    public static synchronized void initWithChannel(ManagedChannel channel) {
        if (instance == null) {
            instance = new RoaureServiceClient(channel);
        }
    }

    public static RoaureServiceClient getInstance() {
        if (instance == null) {
            throw new IllegalStateException(
                    "RoaureServiceClient is not initialized. " +
                            "Call init(host, port) or initWithChannel(channel) first.");
        }
        return instance;
    }

    public void shutdown() {
        if (channel != null) {
            channel.shutdown();
        }
        instance = null;
    }

    private <T> StreamObserver<T> universalStreamObserver(
            OnValueListener<T> onValueListener,
            OnStatusListener onStatusListener,
            OnCompletedListener onCompletedListener
    ) {
        return new StreamObserver<T>() {
            @Override
            public void onNext(T value) {
                onValueListener.onValue(value);
            }

            @Override
            public void onError(Throwable t) {
                Status status = Status.fromThrowable(t);
                onStatusListener.onStatus(status);
            }

            @Override
            public void onCompleted() {
                onCompletedListener.onCompleted();
            }
        };
    }

    public void getMetrics(
            int pollInterval,
            OnValueListener<Metrics> onValueListener,
            OnStatusListener onStatusListener,
            OnCompletedListener onCompletedListener
    ) {
        asyncStub.getMetrics(
                GetMetricsRequest
                        .newBuilder()
                        .setPollInterval(pollInterval)
                        .build(),
                universalStreamObserver(onValueListener, onStatusListener, onCompletedListener)
        );
    }

    public void getMonitorConf(
            OnValueListener<MonitorConf> onValueListener,
            OnStatusListener onStatusListener,
            OnCompletedListener onCompletedListener
    ) {
        asyncStub.getMonitorConf(
                Empty
                        .newBuilder()
                        .build(),
                universalStreamObserver(onValueListener, onStatusListener, onCompletedListener)
        );
    }

    public void getRouterConf(
            OnValueListener<RouterConf> onValueListener,
            OnStatusListener onStatusListener,
            OnCompletedListener onCompletedListener
    ) {
        asyncStub.getRouterConf(
                Empty
                        .newBuilder()
                        .build(),
                universalStreamObserver(onValueListener, onStatusListener, onCompletedListener)
        );
    }

    public void getIperfServerConf(
            OnValueListener<IperfServerConf> onValueListener,
            OnStatusListener onStatusListener,
            OnCompletedListener onCompletedListener
    ) {
        asyncStub.getIperfServerConf(
                Empty
                        .newBuilder()
                        .build(),
                universalStreamObserver(onValueListener, onStatusListener, onCompletedListener)
        );
    }

    public void updateMonitorConf(
            MonitorConf monitorConf,
            OnValueListener<Empty> onValueListener,
            OnStatusListener onStatusListener,
            OnCompletedListener onCompletedListener
    ) {
        asyncStub.updateMonitorConf(
                UpdateMonitorConfRequest
                        .newBuilder()
                        .setMonitorConf(monitorConf)
                        .build(),
                universalStreamObserver(onValueListener, onStatusListener, onCompletedListener)
        );
    }

    public void updateRouterConf(
            RouterConf routerConf,
            OnValueListener<Empty> onValueListener,
            OnStatusListener onStatusListener,
            OnCompletedListener onCompletedListener
    ) {
        asyncStub.updateRouterConf(
                UpdateRouterConfRequest
                        .newBuilder()
                        .setRouterConf(routerConf)
                        .build(),
                universalStreamObserver(onValueListener, onStatusListener, onCompletedListener)
        );
    }

    public void updateIperfServerConf(
            IperfServerConf iperfServerConf,
            OnValueListener<Empty> onValueListener,
            OnStatusListener onStatusListener,
            OnCompletedListener onCompletedListener
    ) {
        asyncStub.updateIperfServerConf(
                UpdateIperfServerConfRequest
                        .newBuilder()
                        .setIperfServerConf(iperfServerConf)
                        .build(),
                universalStreamObserver(onValueListener, onStatusListener, onCompletedListener)
        );
    }

    public void toggleMonitor(
            OnValueListener<Empty> onValueListener,
            OnStatusListener onStatusListener,
            OnCompletedListener onCompletedListener
    ) {
        asyncStub.toggleMonitor(
                Empty
                        .newBuilder()
                        .build(),
                universalStreamObserver(onValueListener, onStatusListener, onCompletedListener)
        );
    }

    public void reboot(
            OnValueListener<Empty> onValueListener,
            OnStatusListener onStatusListener,
            OnCompletedListener onCompletedListener
    ) {
        asyncStub.reboot(
                Empty
                        .newBuilder()
                        .build(),
                universalStreamObserver(onValueListener, onStatusListener, onCompletedListener)
        );
    }

    public void listSchedules(
            OnValueListener<ListSchedulesResponse> onValueListener,
            OnStatusListener onStatusListener,
            OnCompletedListener onCompletedListener
    ) {
        asyncStub.listSchedules(
                Empty
                        .newBuilder()
                        .build(),
                universalStreamObserver(onValueListener, onStatusListener, onCompletedListener)
        );
    }

    public void createSchedule(
            Schedule schedule,
            OnValueListener<Schedule> onValueListener,
            OnStatusListener onStatusListener,
            OnCompletedListener onCompletedListener
    ) {
        asyncStub.createSchedule(
                CreateScheduleRequest
                        .newBuilder()
                        .setSchedule(schedule)
                        .build(),
                universalStreamObserver(onValueListener, onStatusListener, onCompletedListener)
        );
    }

    public void updateSchedule(
            String id,
            Schedule schedule,
            OnValueListener<Schedule> onValueListener,
            OnStatusListener onStatusListener,
            OnCompletedListener onCompletedListener
    ) {
        asyncStub.updateSchedule(
                UpdateScheduleRequest
                        .newBuilder()
                        .setId(id)
                        .setSchedule(schedule)
                        .build(),
                universalStreamObserver(onValueListener, onStatusListener, onCompletedListener)
        );
    }

    public void deleteSchedule(
            String id,
            OnValueListener<Empty> onValueListener,
            OnStatusListener onStatusListener,
            OnCompletedListener onCompletedListener
    ) {
        asyncStub.deleteSchedule(
                DeleteScheduleRequest
                        .newBuilder()
                        .setId(id)
                        .build(),
                universalStreamObserver(onValueListener, onStatusListener, onCompletedListener)
        );
    }
}
