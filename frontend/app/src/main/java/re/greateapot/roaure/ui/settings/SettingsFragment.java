package re.greateapot.roaure.ui.settings;

import android.content.Context;
import android.content.Intent;
import android.os.Bundle;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;

import androidx.annotation.NonNull;
import androidx.annotation.Nullable;
import androidx.fragment.app.Fragment;
import androidx.lifecycle.ViewModelProvider;

import com.google.android.material.snackbar.Snackbar;
import com.google.android.material.textfield.TextInputLayout;

import re.greateapot.roaure.R;
import re.greateapot.roaure.ui.main.MainActivity;

public class SettingsFragment extends Fragment {

    private SettingsViewModel mViewModel;

    @Override
    public View onCreateView(@NonNull LayoutInflater inflater,
                             @Nullable ViewGroup container,
                             @Nullable Bundle savedInstanceState) {
        return inflater.inflate(R.layout.fragment_settings, container, false);
    }

    @Override
    public void onViewCreated(@NonNull View view, @Nullable Bundle savedInstanceState) {
        super.onViewCreated(view, savedInstanceState);
        mViewModel = new ViewModelProvider(this).get(SettingsViewModel.class);

        TextInputLayout h0 = view.findViewById(R.id.backend_host);
        TextInputLayout p0 = view.findViewById(R.id.backend_port);

        view.findViewById(R.id.reset_button).setOnClickListener(view1 -> loadParams());
        view.findViewById(R.id.save_button).setOnClickListener(view1 -> {
            var h1 = h0.getEditText();
            var h2 = h1 != null ? h1.getText() : null;
            var host = h2 != null ? h2.toString() : "";

            var p1 = p0 != null ? p0.getEditText() : null;
            var p2 = p1 != null ? p1.getText() : null;
            var port = Integer.parseInt(p2 != null ? p2.toString() : "0");

            dumpParams(view, host, port);
        });

        mViewModel.getHostValue().observe(getViewLifecycleOwner(), value -> {
            var h1 = h0.getEditText();
            if (h1 != null) h1.setText(value);
        });

        mViewModel.getPortValue().observe(getViewLifecycleOwner(), value -> {
            var p1 = p0.getEditText();
            if (p1 != null) p1.setText(String.valueOf(value));
        });

        loadParams();
    }

    protected void loadParams() {
        var activity = getActivity();
        if (activity == null) return;

        var prefs = getActivity().getSharedPreferences("conf", Context.MODE_PRIVATE);
        var host = prefs.getString(getResources().getString(R.string.backend_host), getResources().getString(R.string.backend_host_default));
        var port = prefs.getInt(getResources().getString(R.string.backend_port), getResources().getInteger(R.integer.backend_port_default));

        mViewModel.getHostValue().postValue(host);
        mViewModel.getPortValue().postValue(port);
    }

    protected void dumpParams(View view, String host, int port) {
        var activity = getActivity();
        if (activity == null) return;

        var prefs = getActivity().getSharedPreferences("conf", Context.MODE_PRIVATE);
        var editor = prefs.edit();

        editor.putString(getResources().getString(R.string.backend_host), host);
        editor.putInt(getResources().getString(R.string.backend_port), port);
        var result = editor.commit();

        if (!result) {
            Snackbar
                    .make(view, "Unable to save", Snackbar.LENGTH_INDEFINITE)
                    .setAction("Retry", view2 -> dumpParams(view, host, port))
                    .show();
            return;
        }

        var intent = new Intent(activity, MainActivity.class);
        intent.addFlags(Intent.FLAG_ACTIVITY_NEW_TASK);
        activity.startActivity(intent);
        activity.finish();
        Runtime.getRuntime().exit(0);
    }
}