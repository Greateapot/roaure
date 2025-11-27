package re.greateapot.roaure;

import android.app.Application;

import re.greateapot.roaure.api.RoaureServiceClient;

public class App extends Application {
    @Override
    public void onCreate() {
        super.onCreate();

        // TODO: load host & port from prefs (with defaults: host: "orangepi3", port: 50052)
        RoaureServiceClient.init("192.168.10.151", 50052);
    }
}
