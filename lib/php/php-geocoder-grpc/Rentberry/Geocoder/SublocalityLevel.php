<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: geocoder.proto

namespace Rentberry\Geocoder;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>SublocalityLevel</code>
 */
class SublocalityLevel extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>int32 level = 1;</code>
     */
    protected $level = 0;
    /**
     * Generated from protobuf field <code>string name = 2;</code>
     */
    protected $name = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type int $level
     *     @type string $name
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Geocoder::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>int32 level = 1;</code>
     * @return int
     */
    public function getLevel()
    {
        return $this->level;
    }

    /**
     * Generated from protobuf field <code>int32 level = 1;</code>
     * @param int $var
     * @return $this
     */
    public function setLevel($var)
    {
        GPBUtil::checkInt32($var);
        $this->level = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string name = 2;</code>
     * @return string
     */
    public function getName()
    {
        return $this->name;
    }

    /**
     * Generated from protobuf field <code>string name = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setName($var)
    {
        GPBUtil::checkString($var, True);
        $this->name = $var;

        return $this;
    }

}

