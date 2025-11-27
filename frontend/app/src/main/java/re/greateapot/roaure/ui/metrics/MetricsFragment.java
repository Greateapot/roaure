package re.greateapot.roaure.ui.metrics;

import android.os.Bundle;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;

import androidx.fragment.app.Fragment;
import androidx.annotation.NonNull;
import androidx.annotation.Nullable;
import androidx.lifecycle.ViewModelProvider;

import re.greateapot.roaure.R;

public class MetricsFragment extends Fragment {
    private re.greateapot.roaure.databinding.FragmentMetricsBinding binding;

    @Nullable
    @Override
    public View onCreateView(@NonNull LayoutInflater inflater,
                             @Nullable ViewGroup container,
                             @Nullable Bundle savedInstanceState) {
        binding = re.greateapot.roaure.databinding.FragmentMetricsBinding.inflate(inflater, container, false);

        // setup viewModel
        binding.setViewModel(new ViewModelProvider(requireActivity()).get(MetricsViewModel.class));
        binding.setLifecycleOwner(getViewLifecycleOwner());

        // bind toggle button logic
        binding.getRoot()
                .findViewById(R.id.toggle_monitor_button)
                .setOnClickListener(view -> binding.getViewModel().toggleMonitor());

        // start polling
        binding.getViewModel().getMetrics();

        return binding.getRoot();
    }

    @Override
    public void onDestroyView() {
        super.onDestroyView();
        binding = null;
    }
}