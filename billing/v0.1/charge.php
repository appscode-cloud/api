<?php
// DO NOT EDIT! Generated by Protobuf-PHP protoc plugin 1.0
// Source: charge.proto
//   Date: 2016-04-20 05:49:21

namespace billing {

  class ChargeCalculateRequest extends \DrSlump\Protobuf\Message {

    /**  @var string */
    public $charge_type = null;
    

    /** @var \Closure[] */
    protected static $__extensions = array();

    public static function descriptor()
    {
      $descriptor = new \DrSlump\Protobuf\Descriptor(__CLASS__, 'billing.ChargeCalculateRequest');

      // OPTIONAL STRING charge_type = 1
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 1;
      $f->name      = "charge_type";
      $f->type      = \DrSlump\Protobuf::TYPE_STRING;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $descriptor->addField($f);

      foreach (self::$__extensions as $cb) {
        $descriptor->addField($cb(), true);
      }

      return $descriptor;
    }

    /**
     * Check if <charge_type> has a value
     *
     * @return boolean
     */
    public function hasChargeType(){
      return $this->_has(1);
    }
    
    /**
     * Clear <charge_type> value
     *
     * @return \billing\ChargeCalculateRequest
     */
    public function clearChargeType(){
      return $this->_clear(1);
    }
    
    /**
     * Get <charge_type> value
     *
     * @return string
     */
    public function getChargeType(){
      return $this->_get(1);
    }
    
    /**
     * Set <charge_type> value
     *
     * @param string $value
     * @return \billing\ChargeCalculateRequest
     */
    public function setChargeType( $value){
      return $this->_set(1, $value);
    }
  }
}

namespace billing {

  class ChargeCalculateResponse extends \DrSlump\Protobuf\Message {

    /**  @var \dtypes\Status */
    public $status = null;
    
    /**  @var string */
    public $cost_usd = null;
    

    /** @var \Closure[] */
    protected static $__extensions = array();

    public static function descriptor()
    {
      $descriptor = new \DrSlump\Protobuf\Descriptor(__CLASS__, 'billing.ChargeCalculateResponse');

      // OPTIONAL MESSAGE status = 1
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 1;
      $f->name      = "status";
      $f->type      = \DrSlump\Protobuf::TYPE_MESSAGE;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $f->reference = '\dtypes\Status';
      $descriptor->addField($f);

      // OPTIONAL STRING cost_usd = 2
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 2;
      $f->name      = "cost_usd";
      $f->type      = \DrSlump\Protobuf::TYPE_STRING;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $descriptor->addField($f);

      foreach (self::$__extensions as $cb) {
        $descriptor->addField($cb(), true);
      }

      return $descriptor;
    }

    /**
     * Check if <status> has a value
     *
     * @return boolean
     */
    public function hasStatus(){
      return $this->_has(1);
    }
    
    /**
     * Clear <status> value
     *
     * @return \billing\ChargeCalculateResponse
     */
    public function clearStatus(){
      return $this->_clear(1);
    }
    
    /**
     * Get <status> value
     *
     * @return \dtypes\Status
     */
    public function getStatus(){
      return $this->_get(1);
    }
    
    /**
     * Set <status> value
     *
     * @param \dtypes\Status $value
     * @return \billing\ChargeCalculateResponse
     */
    public function setStatus(\dtypes\Status $value){
      return $this->_set(1, $value);
    }
    
    /**
     * Check if <cost_usd> has a value
     *
     * @return boolean
     */
    public function hasCostUsd(){
      return $this->_has(2);
    }
    
    /**
     * Clear <cost_usd> value
     *
     * @return \billing\ChargeCalculateResponse
     */
    public function clearCostUsd(){
      return $this->_clear(2);
    }
    
    /**
     * Get <cost_usd> value
     *
     * @return string
     */
    public function getCostUsd(){
      return $this->_get(2);
    }
    
    /**
     * Set <cost_usd> value
     *
     * @param string $value
     * @return \billing\ChargeCalculateResponse
     */
    public function setCostUsd( $value){
      return $this->_set(2, $value);
    }
  }
}

namespace billing {

  class ChargeClient extends \Grpc\BaseStub {

    public function __construct($hostname, $opts) {
      parent::__construct($hostname, $opts);
    }
    /**
     * @param billing\ChargeCalculateRequest $input
     */
    public function Calculate(\billing\ChargeCalculateRequest $argument, $metadata = array(), $options = array()) {
      return $this->_simpleRequest('/billing.Charge/Calculate', $argument, '\billing\ChargeCalculateResponse::deserialize', $metadata, $options);
    }
  }
}
