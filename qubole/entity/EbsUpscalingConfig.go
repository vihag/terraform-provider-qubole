package entity

type EbsUpscalingConfig struct {
	max_ebs_volume_count          int
	percent_free_space_threshold  int
	absolute_free_space_threshold int
	sampling_interval             int
	sampling_window               int
}
