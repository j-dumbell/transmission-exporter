package exporter

import (
	"context"
	"fmt"
	"log/slog"
	"testing"

	"github.com/j-dumbell/go-qbittorrent/pkg/transmission"
	"github.com/prometheus/client_golang/prometheus"
	promclient "github.com/prometheus/client_model/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExporter(t *testing.T) {
	t.Run("exportTorrentLevelMetrics disabled", func(t *testing.T) {
		reg := prometheus.NewRegistry()
		exporter := New(&TestTransmissionClient{}, slog.Default(), false)
		err := reg.Register(exporter)
		require.NoError(t, err, "Register should not error")

		mfs, err := reg.Gather()
		require.NoError(t, err, "Gather should not error")

		assertGlobalMetrics(t, mfs)

		for _, descCfg := range torrentLevelDescConfigs {
			assertMetricFamilyDoesNotExist(t, mfs, string(descCfg.Metric))
		}
	})

	t.Run("exportTorrentLevelMetrics enabled", func(t *testing.T) {
		reg := prometheus.NewRegistry()
		exporter := New(&TestTransmissionClient{}, slog.Default(), true)
		err := reg.Register(exporter)
		require.NoError(t, err, "Register should not error")

		mfs, err := reg.Gather()
		require.NoError(t, err, "Gather should not error")

		assertGlobalMetrics(t, mfs)

		assertMetricValueWithLabels(t, mfs, metricNameTorrentDownloadBytesPerSecond, prometheus.GaugeValue, []MetricValue{
			{Labels: map[string]string{hashLabel: t1.HashString}, Value: float64(t1.RateDownload)},
			{Labels: map[string]string{hashLabel: t2.HashString}, Value: float64(t2.RateDownload)},
		})
		assertMetricValueWithLabels(t, mfs, metricNameTorrentUploadBytesPerSecond, prometheus.GaugeValue, []MetricValue{
			{Labels: map[string]string{hashLabel: t1.HashString}, Value: float64(t1.RateUpload)},
			{Labels: map[string]string{hashLabel: t2.HashString}, Value: float64(t2.RateUpload)},
		})
		assertMetricValueWithLabels(t, mfs, metricNameTorrentTotalSizeBytes, prometheus.GaugeValue, []MetricValue{
			{Labels: map[string]string{hashLabel: t1.HashString}, Value: float64(t1.TotalSize)},
			{Labels: map[string]string{hashLabel: t2.HashString}, Value: float64(t2.TotalSize)},
		})
		assertMetricValueWithLabels(t, mfs, metricNameTorrentSizeWhenDoneBytes, prometheus.GaugeValue, []MetricValue{
			{Labels: map[string]string{hashLabel: t1.HashString}, Value: float64(t1.SizeWhenDone)},
			{Labels: map[string]string{hashLabel: t2.HashString}, Value: float64(t2.SizeWhenDone)},
		})
		assertMetricValueWithLabels(t, mfs, metricNameTorrentLeftUntilDoneBytes, prometheus.GaugeValue, []MetricValue{
			{Labels: map[string]string{hashLabel: t1.HashString}, Value: float64(t1.LeftUntilDone)},
			{Labels: map[string]string{hashLabel: t2.HashString}, Value: float64(t2.LeftUntilDone)},
		})
		assertMetricValueWithLabels(t, mfs, metricNameTorrentDownloadBytesTotal, prometheus.CounterValue, []MetricValue{
			{Labels: map[string]string{hashLabel: t1.HashString}, Value: float64(t1.DownloadedEver)},
			{Labels: map[string]string{hashLabel: t2.HashString}, Value: float64(t2.DownloadedEver)},
		})
		assertMetricValueWithLabels(t, mfs, metricNameTorrentUploadBytesTotal, prometheus.CounterValue, []MetricValue{
			{Labels: map[string]string{hashLabel: t1.HashString}, Value: float64(t1.UploadedEver)},
			{Labels: map[string]string{hashLabel: t2.HashString}, Value: float64(t2.UploadedEver)},
		})
		assertMetricValueWithLabels(t, mfs, metricNameTorrentCorruptBytesTotal, prometheus.CounterValue, []MetricValue{
			{Labels: map[string]string{hashLabel: t1.HashString}, Value: float64(t1.CorruptEver)},
			{Labels: map[string]string{hashLabel: t2.HashString}, Value: float64(t2.CorruptEver)},
		})
		assertMetricValueWithLabels(t, mfs, metricNameTorrentPeersConnected, prometheus.GaugeValue, []MetricValue{
			{Labels: map[string]string{hashLabel: t1.HashString}, Value: float64(t1.PeersConnected)},
			{Labels: map[string]string{hashLabel: t2.HashString}, Value: float64(t2.PeersConnected)},
		})
		assertMetricValueWithLabels(t, mfs, metricNameTorrentPeersSendingToUs, prometheus.GaugeValue, []MetricValue{
			{Labels: map[string]string{hashLabel: t1.HashString}, Value: float64(t1.PeersSendingToUs)},
			{Labels: map[string]string{hashLabel: t2.HashString}, Value: float64(t2.PeersSendingToUs)},
		})
		assertMetricValueWithLabels(t, mfs, metricNameTorrentPeersGettingFromUs, prometheus.GaugeValue, []MetricValue{
			{Labels: map[string]string{hashLabel: t1.HashString}, Value: float64(t1.PeersGettingFromUs)},
			{Labels: map[string]string{hashLabel: t2.HashString}, Value: float64(t2.PeersGettingFromUs)},
		})
		assertMetricValueWithLabels(t, mfs, metricNameTorrentWebseedsSendingToUs, prometheus.GaugeValue, []MetricValue{
			{Labels: map[string]string{hashLabel: t1.HashString}, Value: float64(t1.WebseedsSendingToUs)},
			{Labels: map[string]string{hashLabel: t2.HashString}, Value: float64(t2.WebseedsSendingToUs)},
		})
		assertMetricValueWithLabels(t, mfs, metricNameTorrentSecondsDownloadingTotal, prometheus.CounterValue, []MetricValue{
			{Labels: map[string]string{hashLabel: t1.HashString}, Value: float64(t1.SecondsDownloading)},
			{Labels: map[string]string{hashLabel: t2.HashString}, Value: float64(t2.SecondsDownloading)},
		})
		assertMetricValueWithLabels(t, mfs, metricNameTorrentSecondsSeedingTotal, prometheus.CounterValue, []MetricValue{
			{Labels: map[string]string{hashLabel: t1.HashString}, Value: float64(t1.SecondsSeeding)},
			{Labels: map[string]string{hashLabel: t2.HashString}, Value: float64(t2.SecondsSeeding)},
		})
		assertMetricValueWithLabels(t, mfs, metricNameTorrentStatus, prometheus.GaugeValue, []MetricValue{
			{Labels: map[string]string{hashLabel: t1.HashString}, Value: float64(t1.Status)},
			{Labels: map[string]string{hashLabel: t2.HashString}, Value: float64(t2.Status)},
		})
		assertMetricValueWithLabels(t, mfs, metricNameTorrentInfo, prometheus.GaugeValue, []MetricValue{
			{Labels: map[string]string{hashLabel: t1.HashString, nameLabel: t1.Name}, Value: float64(1)},
			{Labels: map[string]string{hashLabel: t2.HashString, nameLabel: t2.Name}, Value: float64(1)},
		})
	})
}

func assertGlobalMetrics(t *testing.T, mfs []*promclient.MetricFamily) {
	assertMetricValue(t, mfs, metricNameDownloadedBytesTotal, prometheus.CounterValue, float64(mockSessionStatsResult.CumulativeStats.DownloadedBytes))
	assertMetricValue(t, mfs, metricNameUploadedBytesTotal, prometheus.CounterValue, float64(mockSessionStatsResult.CumulativeStats.UploadedBytes))
	assertMetricValue(t, mfs, metricNameTorrentsAddedTotal, prometheus.CounterValue, float64(mockSessionStatsResult.CumulativeStats.FilesAdded))
	assertMetricValue(t, mfs, metricNameSecondsActiveTotal, prometheus.CounterValue, float64(mockSessionStatsResult.CumulativeStats.SecondsActive))
	assertMetricValue(t, mfs, metricNameSessionsTotal, prometheus.CounterValue, float64(mockSessionStatsResult.CumulativeStats.SessionCount))
	assertMetricValue(t, mfs, metricNameUploadBytesPerSecond, prometheus.GaugeValue, float64(mockSessionStatsResult.UploadSpeed))
	assertMetricValue(t, mfs, metricNameDownloadBytesPerSecond, prometheus.GaugeValue, float64(mockSessionStatsResult.DownloadSpeed))
	assertMetricValueWithLabels(t, mfs, metricNameTorrents, prometheus.GaugeValue, []MetricValue{
		{Labels: map[string]string{statusLabel: transmission.TorrentStatusStopped.String()}, Value: 0},
		{Labels: map[string]string{statusLabel: transmission.TorrentStatusCheckWait.String()}, Value: 0},
		{Labels: map[string]string{statusLabel: transmission.TorrentStatusCheck.String()}, Value: 0},
		{Labels: map[string]string{statusLabel: transmission.TorrentStatusDownloadWait.String()}, Value: 0},
		{Labels: map[string]string{statusLabel: transmission.TorrentStatusDownload.String()}, Value: 1},
		{Labels: map[string]string{statusLabel: transmission.TorrentStatusSeedWait.String()}, Value: 0},
		{Labels: map[string]string{statusLabel: transmission.TorrentStatusSeed.String()}, Value: 1},
		{Labels: map[string]string{statusLabel: "unknown"}, Value: 0},
	})
	assertMetricValueWithLabels(t, mfs, metricNameVersion, prometheus.GaugeValue, []MetricValue{
		{Labels: map[string]string{versionLabel: mockSession.Version}, Value: 1},
	})
}

type MetricValue struct {
	Labels map[string]string
	Value  float64
}

func assertMetricValueWithLabels(t *testing.T, mfs []*promclient.MetricFamily, metricName metricName, valueType prometheus.ValueType, expectedMetricValues []MetricValue) {
	mf := findMetricFamily(mfs, string(metricName))
	if mf == nil {
		assert.Fail(t, "metric family not found", metricName)
		return
	}

	var actual []MetricValue
	for _, metric := range mf.GetMetric() {
		labels := make(map[string]string)
		for _, labelPair := range metric.GetLabel() {
			labels[labelPair.GetName()] = labelPair.GetValue()
		}
		actual = append(actual, MetricValue{
			Labels: labels,
			Value:  getValue(metric, valueType),
		})
	}

	assert.ElementsMatchf(t, expectedMetricValues, actual, "unexpected metric values for metric %s", metricName)
}

func assertMetricValue(t *testing.T, mfs []*promclient.MetricFamily, metricName metricName, valueType prometheus.ValueType, expectedValue float64) {
	mf := findMetricFamily(mfs, string(metricName))
	if mf == nil {
		assert.Fail(t, "metric family not found", metricName)
		return
	}

	assert.NotEqualf(t, "", mf.GetHelp(), "metric %s should have help text", metricName)

	metrics := mf.GetMetric()
	if len(metrics) != 1 {
		assert.Failf(t, "found more than 1 metric", "metric name: %s; found: %d", metricName, len(metrics))
		return
	}

	actual := getValue(metrics[0], valueType)

	assert.Equalf(t, expectedValue, actual, "unexpected metric value for %s", metricName)
}

func getValue(metric *promclient.Metric, valueType prometheus.ValueType) float64 {
	switch valueType {
	case prometheus.CounterValue:
		return metric.GetCounter().GetValue()
	case prometheus.GaugeValue:
		return metric.GetGauge().GetValue()
	case prometheus.UntypedValue:
		return metric.GetUntyped().GetValue()
	}
	panic(fmt.Sprint("invalid metric valueType: ", valueType))
}

func findMetricFamily(mfs []*promclient.MetricFamily, name string) *promclient.MetricFamily {
	for _, mf := range mfs {
		if mf.GetName() == name {
			return mf
		}
	}
	return nil
}

func assertMetricFamilyDoesNotExist(t *testing.T, mfs []*promclient.MetricFamily, name string) {
	t.Helper()
	for _, mf := range mfs {
		if mf.GetName() == name {
			assert.Fail(t, "found unexpected metric family", name)
		}
	}
}

type TestTransmissionClient struct {
}

func (t *TestTransmissionClient) SessionStats(_ context.Context) (*transmission.SessionStatsResult, error) {
	return &mockSessionStatsResult, nil
}

func (t *TestTransmissionClient) SessionGet(_ context.Context) (*transmission.Session, error) {
	return &mockSession, nil
}

func (t *TestTransmissionClient) TorrentGet(_ context.Context, _ transmission.TorrentGetArgs) (*transmission.TorrentGetResult, error) {
	torrentGetResult := transmission.TorrentGetResult{
		Torrents: []transmission.Torrent{t1, t2},
		Removed:  []int64{},
	}
	return &torrentGetResult, nil
}

var mockSessionStatsResult = transmission.SessionStatsResult{
	ActiveTorrentCount: 1,
	DownloadSpeed:      123,
	PausedTorrentCount: 2,
	TorrentCount:       3,
	UploadSpeed:        456,
	CumulativeStats: transmission.Stats{
		UploadedBytes:   321,
		DownloadedBytes: 654,
		FilesAdded:      3,
		SecondsActive:   99,
		SessionCount:    22,
	},
}

var mockSession = transmission.Session{
	Version: "4.0.0",
}

var t1 = transmission.Torrent{
	HashString:          "abc",
	Name:                "foo",
	Status:              transmission.TorrentStatusDownload,
	RateDownload:        11,
	RateUpload:          22,
	TotalSize:           33,
	SizeWhenDone:        44,
	LeftUntilDone:       55,
	DownloadedEver:      66,
	UploadedEver:        77,
	CorruptEver:         88,
	PeersConnected:      99,
	PeersSendingToUs:    111,
	PeersGettingFromUs:  222,
	WebseedsSendingToUs: 333,
	SecondsDownloading:  444,
	SecondsSeeding:      555,
}

var t2 = transmission.Torrent{
	HashString:          "def",
	Name:                "bar",
	Status:              transmission.TorrentStatusSeed,
	RateDownload:        1111,
	RateUpload:          2222,
	TotalSize:           3333,
	SizeWhenDone:        4444,
	LeftUntilDone:       5555,
	DownloadedEver:      6666,
	UploadedEver:        7777,
	CorruptEver:         8888,
	PeersConnected:      9999,
	PeersSendingToUs:    11111,
	PeersGettingFromUs:  22222,
	WebseedsSendingToUs: 33333,
	SecondsDownloading:  44444,
	SecondsSeeding:      55555,
}
