package re.greateapot.roaure;

import android.app.Application;

import re.greateapot.roaure.api.RoaureServiceClient;

public class App extends Application {
    @Override
    public void onCreate() {
        super.onCreate();

        var prefs = getSharedPreferences("conf", MODE_PRIVATE);
        var host = prefs.getString(getResources().getString(R.string.backend_host), getResources().getString(R.string.backend_host_default));
        var port = prefs.getInt(getResources().getString(R.string.backend_port), getResources().getInteger(R.integer.backend_port_default));

        RoaureServiceClient.init(host, port);
    }
}
