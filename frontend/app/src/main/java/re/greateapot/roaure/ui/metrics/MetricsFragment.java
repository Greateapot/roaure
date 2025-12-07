package re.greateapot.roaure.ui.metrics;

import android.os.Bundle;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.TextView;

import androidx.annotation.NonNull;
import androidx.annotation.Nullable;
import androidx.fragment.app.Fragment;
import androidx.fragment.app.FragmentActivity;
import androidx.lifecycle.ViewModelProvider;

import com.google.android.material.button.MaterialButton;

import java.util.Locale;

import re.greateapot.roaure.R;

public class MetricsFragment extends Fragment {

    private MetricsViewModel mViewModel;

    @Nullable
    @Override
    public View onCreateView(@NonNull LayoutInflater inflater,
                             @Nullable ViewGroup container,
                             @Nullable Bundle savedInstanceState) {
        return inflater.inflate(R.layout.fragment_metrics, container, false);
    }

    @Override
    public void onViewCreated(@NonNull View view, @Nullable Bundle savedInstanceState) {
        super.onViewCreated(view, savedInstanceState);
        mViewModel = new ViewModelProvider(this).get(MetricsViewModel.class);


        view.findViewById(R.id.toggle_monitor_button).setOnClickListener(view1 -> mViewModel.toggleMonitor());

        mViewModel.getDownloadSpeedValue().observe(getViewLifecycleOwner(), value -> {
            FragmentActivity activity = getActivity();
            if (activity == null) return;

            activity.runOnUiThread(() -> {
                TextView tv = view.findViewById(R.id.download_speed_text_view);
                if (value == null) {
                    tv.setText(R.string.no_value_present);
                } else {
                    // TODO: RW
                    tv.setText(String.format(Locale.getDefault(), "%.2f", value / 1024 / 1024));
                }
            });
        });

        mViewModel.getBadCountValue().observe(getViewLifecycleOwner(), value -> {
            FragmentActivity activity = getActivity();
            if (activity == null) return;

            activity.runOnUiThread(() -> {
                TextView tv = view.findViewById(R.id.bad_count_text_view);
                if (value == null) {
                    tv.setText(R.string.no_value_present);
                } else {
                    tv.setText(String.valueOf(value));
                }
            });
        });

        mViewModel.getRebootRequiredValue().observe(getViewLifecycleOwner(), value -> {
            FragmentActivity activity = getActivity();
            if (activity == null) return;

            activity.runOnUiThread(() -> {
                TextView tv = view.findViewById(R.id.reboot_required_text_view);
                if (value == null) {
                    tv.setText(R.string.no_value_present);
                } else {
                    // TODO: RW
                    tv.setText(value ? "YES" : "NO");
                }
            });
        });

        mViewModel.getMonitorRunningValue().observe(getViewLifecycleOwner(), value -> {
            FragmentActivity activity = getActivity();
            if (activity == null) return;

            activity.runOnUiThread(() -> {
                MaterialButton mb = view.findViewById(R.id.toggle_monitor_button);
                if (value == null) {
                    mb.setText(R.string.no_value_present);
                    mb.setSelected(false);
                } else if (value) {
                    mb.setText(R.string.stop_monitor);
                    mb.setSelected(true);
                } else {
                    mb.setText(R.string.start_monitor);
                    mb.setSelected(false);
                }
            });
        });

        // Start polling
        mViewModel.getMetrics();
    }
}