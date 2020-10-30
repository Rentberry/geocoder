<?php
// GENERATED CODE -- DO NOT EDIT!

namespace Rentberry\Geocoder;

/**
 */
class TimezoneServiceClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * @param \Rentberry\Geocoder\TimezoneRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function Lookup(\Rentberry\Geocoder\TimezoneRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/TimezoneService/Lookup',
        $argument,
        ['\Rentberry\Geocoder\Timezone', 'decode'],
        $metadata, $options);
    }

}
