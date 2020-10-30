<?php
// GENERATED CODE -- DO NOT EDIT!

namespace Rentberry\Geocoder;

/**
 */
class GeocodeServiceClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * @param \Rentberry\Geocoder\LocationRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function Geocode(\Rentberry\Geocoder\LocationRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/GeocodeService/Geocode',
        $argument,
        ['\Rentberry\Geocoder\LocationResponse', 'decode'],
        $metadata, $options);
    }

}
