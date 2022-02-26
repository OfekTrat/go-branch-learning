package backtester

import (
	candle_stream "branch_learning/candle_stream"
)

func (b *BackTester) TestMultipleStreams(streams []*candle_stream.CandleStream) {
	for _, stream := range streams {
		b.Test(stream)
	}
}

func TestMultipleBacktesters(backtesters []*BackTester, streams []*candle_stream.CandleStream) {
	for _, backtester := range backtesters {
		backtester.TestMultipleStreams(streams)
	}
}
