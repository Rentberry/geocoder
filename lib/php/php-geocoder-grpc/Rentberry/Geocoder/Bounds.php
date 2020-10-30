<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: geocoder.proto

namespace Rentberry\Geocoder;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>Bounds</code>
 */
class Bounds extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>.LatLng northEast = 1;</code>
     */
    protected $northEast = null;
    /**
     * Generated from protobuf field <code>.LatLng southWest = 2;</code>
     */
    protected $southWest = null;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Rentberry\Geocoder\LatLng $northEast
     *     @type \Rentberry\Geocoder\LatLng $southWest
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Geocoder::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>.LatLng northEast = 1;</code>
     * @return \Rentberry\Geocoder\LatLng
     */
    public function getNorthEast()
    {
        return isset($this->northEast) ? $this->northEast : null;
    }

    public function hasNorthEast()
    {
        return isset($this->northEast);
    }

    public function clearNorthEast()
    {
        unset($this->northEast);
    }

    /**
     * Generated from protobuf field <code>.LatLng northEast = 1;</code>
     * @param \Rentberry\Geocoder\LatLng $var
     * @return $this
     */
    public function setNorthEast($var)
    {
        GPBUtil::checkMessage($var, \Rentberry\Geocoder\LatLng::class);
        $this->northEast = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>.LatLng southWest = 2;</code>
     * @return \Rentberry\Geocoder\LatLng
     */
    public function getSouthWest()
    {
        return isset($this->southWest) ? $this->southWest : null;
    }

    public function hasSouthWest()
    {
        return isset($this->southWest);
    }

    public function clearSouthWest()
    {
        unset($this->southWest);
    }

    /**
     * Generated from protobuf field <code>.LatLng southWest = 2;</code>
     * @param \Rentberry\Geocoder\LatLng $var
     * @return $this
     */
    public function setSouthWest($var)
    {
        GPBUtil::checkMessage($var, \Rentberry\Geocoder\LatLng::class);
        $this->southWest = $var;

        return $this;
    }

}

