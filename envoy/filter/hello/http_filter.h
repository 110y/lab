#pragma once

#include <string>

#include "envoy/filter/hello/http_filter.pb.h"
#include "envoy/server/filter_config.h"

namespace Envoy {
namespace Http {

class HttpSampleDecoderFilterConfig {
   public:
    HttpSampleDecoderFilterConfig(const sample::Decoder& proto_config);

    const std::string& key() const { return key_; }
    const std::string& val() const { return val_; }

   private:
    const std::string key_;
    const std::string val_;
};

typedef std::shared_ptr<HttpSampleDecoderFilterConfig>
    HttpSampleDecoderFilterConfigSharedPtr;

class HttpSampleDecoderFilter : public StreamDecoderFilter {
   public:
    HttpSampleDecoderFilter(HttpSampleDecoderFilterConfigSharedPtr);
    ~HttpSampleDecoderFilter();

    // Http::StreamFilterBase
    void onDestroy() override;

    // Http::StreamDecoderFilter
    FilterHeadersStatus decodeHeaders(HeaderMap&, bool) override;
    FilterDataStatus decodeData(Buffer::Instance&, bool) override;
    FilterTrailersStatus decodeTrailers(HeaderMap&) override;
    void setDecoderFilterCallbacks(StreamDecoderFilterCallbacks&) override;

   private:
    const HttpSampleDecoderFilterConfigSharedPtr config_;
    StreamDecoderFilterCallbacks* decoder_callbacks_;

    const LowerCaseString headerKey() const;
    const std::string headerValue() const;
};

}  // namespace Http
}  // namespace Envoy
