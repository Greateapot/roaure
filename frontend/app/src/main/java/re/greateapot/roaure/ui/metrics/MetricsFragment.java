package re.greateapot.roaure.ui.metrics;

import android.os.Bundle;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;

import androidx.fragment.app.Fragment;
import androidx.annotation.NonNull;
import androidx.annotation.Nullable;
import androidx.lifecycle.ViewModelProvider;

public class MetricsFragment extends Fragment {
    private re.greateapot.roaure.databinding.FragmentMetricsBinding binding;

    @Nullable
    @Override
    public View onCreateView(@NonNull LayoutInflater inflater,
                             @Nullable ViewGroup container,
                             @Nullable Bundle savedInstanceState) {
        binding = re.greateapot.roaure.databinding.FragmentMetricsBinding.inflate(inflater, container, false);
        binding.setViewModel(new ViewModelProvider(requireActivity()).get(MetricsViewModel.class));
        binding.setLifecycleOwner(getViewLifecycleOwner());
        binding.getViewModel().getMetrics();
        return binding.getRoot();
    }

    @Override
    public void onDestroyView() {
        super.onDestroyView();
        binding = null;
    }
}