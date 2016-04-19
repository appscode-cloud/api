<?php
// DO NOT EDIT! Generated by Protobuf-PHP protoc plugin 1.0
// Source: api/dtypes/types.proto
//   Date: 2016-04-19 16:27:29

namespace dtypes {

  class StatusCode extends \DrSlump\Protobuf\Enum {
    const OK = 0;
    const FAILED = 1;
    const UNAUTHORIZED = 2;
    const BADREQUEST = 3;
    const PERMISSION_DENIED = 4;
    const NOT_FOUND = 5;
    const UNIMPLEMENTED = 6;
    const INTERNAL = 7;
    const EXTERNAL = 8;
    const BAD_RESPONSE = 9;
    const UNKNOWN_ERROR = 10;
    const QUOTA_LIMIT_EXCEED = 11;
    const INVALID_QUOTA = 12;
    const PAYMENT_INFORMATION_UNAVAILABLE = 13;
    const INVALID_PAYMENT_INFORMATION = 14;
    const TRANSACTION_FAILED = 15;
    const ARE_YOU_SURE = 16;
  }
}
namespace dtypes {

  class Status extends \DrSlump\Protobuf\Message {

    /**  @var string */
    public $code = null;
    
    /**  @var string */
    public $info = null;
    
    /**  @var string */
    public $message = null;
    
    /**  @var \dtypes\Help */
    public $help = null;
    
    /**  @var \google\protobuf\Any[]  */
    public $details = array();
    

    /** @var \Closure[] */
    protected static $__extensions = array();

    public static function descriptor()
    {
      $descriptor = new \DrSlump\Protobuf\Descriptor(__CLASS__, 'dtypes.Status');

      // OPTIONAL STRING code = 1
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 1;
      $f->name      = "code";
      $f->type      = \DrSlump\Protobuf::TYPE_STRING;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $descriptor->addField($f);

      // OPTIONAL STRING info = 2
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 2;
      $f->name      = "info";
      $f->type      = \DrSlump\Protobuf::TYPE_STRING;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $descriptor->addField($f);

      // OPTIONAL STRING message = 3
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 3;
      $f->name      = "message";
      $f->type      = \DrSlump\Protobuf::TYPE_STRING;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $descriptor->addField($f);

      // OPTIONAL MESSAGE help = 4
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 4;
      $f->name      = "help";
      $f->type      = \DrSlump\Protobuf::TYPE_MESSAGE;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $f->reference = '\dtypes\Help';
      $descriptor->addField($f);

      // REPEATED MESSAGE details = 5
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 5;
      $f->name      = "details";
      $f->type      = \DrSlump\Protobuf::TYPE_MESSAGE;
      $f->rule      = \DrSlump\Protobuf::RULE_REPEATED;
      $f->reference = '\google\protobuf\Any';
      $descriptor->addField($f);

      foreach (self::$__extensions as $cb) {
        $descriptor->addField($cb(), true);
      }

      return $descriptor;
    }

    /**
     * Check if <code> has a value
     *
     * @return boolean
     */
    public function hasCode(){
      return $this->_has(1);
    }
    
    /**
     * Clear <code> value
     *
     * @return \dtypes\Status
     */
    public function clearCode(){
      return $this->_clear(1);
    }
    
    /**
     * Get <code> value
     *
     * @return string
     */
    public function getCode(){
      return $this->_get(1);
    }
    
    /**
     * Set <code> value
     *
     * @param string $value
     * @return \dtypes\Status
     */
    public function setCode( $value){
      return $this->_set(1, $value);
    }
    
    /**
     * Check if <info> has a value
     *
     * @return boolean
     */
    public function hasInfo(){
      return $this->_has(2);
    }
    
    /**
     * Clear <info> value
     *
     * @return \dtypes\Status
     */
    public function clearInfo(){
      return $this->_clear(2);
    }
    
    /**
     * Get <info> value
     *
     * @return string
     */
    public function getInfo(){
      return $this->_get(2);
    }
    
    /**
     * Set <info> value
     *
     * @param string $value
     * @return \dtypes\Status
     */
    public function setInfo( $value){
      return $this->_set(2, $value);
    }
    
    /**
     * Check if <message> has a value
     *
     * @return boolean
     */
    public function hasMessage(){
      return $this->_has(3);
    }
    
    /**
     * Clear <message> value
     *
     * @return \dtypes\Status
     */
    public function clearMessage(){
      return $this->_clear(3);
    }
    
    /**
     * Get <message> value
     *
     * @return string
     */
    public function getMessage(){
      return $this->_get(3);
    }
    
    /**
     * Set <message> value
     *
     * @param string $value
     * @return \dtypes\Status
     */
    public function setMessage( $value){
      return $this->_set(3, $value);
    }
    
    /**
     * Check if <help> has a value
     *
     * @return boolean
     */
    public function hasHelp(){
      return $this->_has(4);
    }
    
    /**
     * Clear <help> value
     *
     * @return \dtypes\Status
     */
    public function clearHelp(){
      return $this->_clear(4);
    }
    
    /**
     * Get <help> value
     *
     * @return \dtypes\Help
     */
    public function getHelp(){
      return $this->_get(4);
    }
    
    /**
     * Set <help> value
     *
     * @param \dtypes\Help $value
     * @return \dtypes\Status
     */
    public function setHelp(\dtypes\Help $value){
      return $this->_set(4, $value);
    }
    
    /**
     * Check if <details> has a value
     *
     * @return boolean
     */
    public function hasDetails(){
      return $this->_has(5);
    }
    
    /**
     * Clear <details> value
     *
     * @return \dtypes\Status
     */
    public function clearDetails(){
      return $this->_clear(5);
    }
    
    /**
     * Get <details> value
     *
     * @param int $idx
     * @return \google\protobuf\Any
     */
    public function getDetails($idx = NULL){
      return $this->_get(5, $idx);
    }
    
    /**
     * Set <details> value
     *
     * @param \google\protobuf\Any $value
     * @return \dtypes\Status
     */
    public function setDetails(\google\protobuf\Any $value, $idx = NULL){
      return $this->_set(5, $value, $idx);
    }
    
    /**
     * Get all elements of <details>
     *
     * @return \google\protobuf\Any[]
     */
    public function getDetailsList(){
     return $this->_get(5);
    }
    
    /**
     * Add a new element to <details>
     *
     * @param \google\protobuf\Any $value
     * @return \dtypes\Status
     */
    public function addDetails(\google\protobuf\Any $value){
     return $this->_add(5, $value);
    }
  }
}

namespace dtypes {

  class Help extends \DrSlump\Protobuf\Message {

    /**  @var string */
    public $description = null;
    
    /**  @var string */
    public $url = null;
    

    /** @var \Closure[] */
    protected static $__extensions = array();

    public static function descriptor()
    {
      $descriptor = new \DrSlump\Protobuf\Descriptor(__CLASS__, 'dtypes.Help');

      // OPTIONAL STRING description = 1
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 1;
      $f->name      = "description";
      $f->type      = \DrSlump\Protobuf::TYPE_STRING;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $descriptor->addField($f);

      // OPTIONAL STRING url = 2
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 2;
      $f->name      = "url";
      $f->type      = \DrSlump\Protobuf::TYPE_STRING;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $descriptor->addField($f);

      foreach (self::$__extensions as $cb) {
        $descriptor->addField($cb(), true);
      }

      return $descriptor;
    }

    /**
     * Check if <description> has a value
     *
     * @return boolean
     */
    public function hasDescription(){
      return $this->_has(1);
    }
    
    /**
     * Clear <description> value
     *
     * @return \dtypes\Help
     */
    public function clearDescription(){
      return $this->_clear(1);
    }
    
    /**
     * Get <description> value
     *
     * @return string
     */
    public function getDescription(){
      return $this->_get(1);
    }
    
    /**
     * Set <description> value
     *
     * @param string $value
     * @return \dtypes\Help
     */
    public function setDescription( $value){
      return $this->_set(1, $value);
    }
    
    /**
     * Check if <url> has a value
     *
     * @return boolean
     */
    public function hasUrl(){
      return $this->_has(2);
    }
    
    /**
     * Clear <url> value
     *
     * @return \dtypes\Help
     */
    public function clearUrl(){
      return $this->_clear(2);
    }
    
    /**
     * Get <url> value
     *
     * @return string
     */
    public function getUrl(){
      return $this->_get(2);
    }
    
    /**
     * Set <url> value
     *
     * @param string $value
     * @return \dtypes\Help
     */
    public function setUrl( $value){
      return $this->_set(2, $value);
    }
  }
}

namespace dtypes {

  class ErrorDetails extends \DrSlump\Protobuf\Message {

    /**  @var string */
    public $requested_resource = null;
    
    /**  @var string */
    public $stacktrace = null;
    

    /** @var \Closure[] */
    protected static $__extensions = array();

    public static function descriptor()
    {
      $descriptor = new \DrSlump\Protobuf\Descriptor(__CLASS__, 'dtypes.ErrorDetails');

      // OPTIONAL STRING requested_resource = 1
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 1;
      $f->name      = "requested_resource";
      $f->type      = \DrSlump\Protobuf::TYPE_STRING;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $descriptor->addField($f);

      // OPTIONAL STRING stacktrace = 2
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 2;
      $f->name      = "stacktrace";
      $f->type      = \DrSlump\Protobuf::TYPE_STRING;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $descriptor->addField($f);

      foreach (self::$__extensions as $cb) {
        $descriptor->addField($cb(), true);
      }

      return $descriptor;
    }

    /**
     * Check if <requested_resource> has a value
     *
     * @return boolean
     */
    public function hasRequestedResource(){
      return $this->_has(1);
    }
    
    /**
     * Clear <requested_resource> value
     *
     * @return \dtypes\ErrorDetails
     */
    public function clearRequestedResource(){
      return $this->_clear(1);
    }
    
    /**
     * Get <requested_resource> value
     *
     * @return string
     */
    public function getRequestedResource(){
      return $this->_get(1);
    }
    
    /**
     * Set <requested_resource> value
     *
     * @param string $value
     * @return \dtypes\ErrorDetails
     */
    public function setRequestedResource( $value){
      return $this->_set(1, $value);
    }
    
    /**
     * Check if <stacktrace> has a value
     *
     * @return boolean
     */
    public function hasStacktrace(){
      return $this->_has(2);
    }
    
    /**
     * Clear <stacktrace> value
     *
     * @return \dtypes\ErrorDetails
     */
    public function clearStacktrace(){
      return $this->_clear(2);
    }
    
    /**
     * Get <stacktrace> value
     *
     * @return string
     */
    public function getStacktrace(){
      return $this->_get(2);
    }
    
    /**
     * Set <stacktrace> value
     *
     * @param string $value
     * @return \dtypes\ErrorDetails
     */
    public function setStacktrace( $value){
      return $this->_set(2, $value);
    }
  }
}

namespace dtypes {

  class LongRunningResponse extends \DrSlump\Protobuf\Message {

    /**  @var \dtypes\Status */
    public $status = null;
    
    /**  @var string */
    public $type_id = null;
    
    /**  @var string */
    public $job_id = null;
    
    /**  @var string */
    public $job_type = null;
    

    /** @var \Closure[] */
    protected static $__extensions = array();

    public static function descriptor()
    {
      $descriptor = new \DrSlump\Protobuf\Descriptor(__CLASS__, 'dtypes.LongRunningResponse');

      // OPTIONAL MESSAGE status = 1
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 1;
      $f->name      = "status";
      $f->type      = \DrSlump\Protobuf::TYPE_MESSAGE;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $f->reference = '\dtypes\Status';
      $descriptor->addField($f);

      // OPTIONAL STRING type_id = 2
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 2;
      $f->name      = "type_id";
      $f->type      = \DrSlump\Protobuf::TYPE_STRING;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $descriptor->addField($f);

      // OPTIONAL STRING job_id = 3
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 3;
      $f->name      = "job_id";
      $f->type      = \DrSlump\Protobuf::TYPE_STRING;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $descriptor->addField($f);

      // OPTIONAL STRING job_type = 4
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 4;
      $f->name      = "job_type";
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
     * @return \dtypes\LongRunningResponse
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
     * @return \dtypes\LongRunningResponse
     */
    public function setStatus(\dtypes\Status $value){
      return $this->_set(1, $value);
    }
    
    /**
     * Check if <type_id> has a value
     *
     * @return boolean
     */
    public function hasTypeId(){
      return $this->_has(2);
    }
    
    /**
     * Clear <type_id> value
     *
     * @return \dtypes\LongRunningResponse
     */
    public function clearTypeId(){
      return $this->_clear(2);
    }
    
    /**
     * Get <type_id> value
     *
     * @return string
     */
    public function getTypeId(){
      return $this->_get(2);
    }
    
    /**
     * Set <type_id> value
     *
     * @param string $value
     * @return \dtypes\LongRunningResponse
     */
    public function setTypeId( $value){
      return $this->_set(2, $value);
    }
    
    /**
     * Check if <job_id> has a value
     *
     * @return boolean
     */
    public function hasJobId(){
      return $this->_has(3);
    }
    
    /**
     * Clear <job_id> value
     *
     * @return \dtypes\LongRunningResponse
     */
    public function clearJobId(){
      return $this->_clear(3);
    }
    
    /**
     * Get <job_id> value
     *
     * @return string
     */
    public function getJobId(){
      return $this->_get(3);
    }
    
    /**
     * Set <job_id> value
     *
     * @param string $value
     * @return \dtypes\LongRunningResponse
     */
    public function setJobId( $value){
      return $this->_set(3, $value);
    }
    
    /**
     * Check if <job_type> has a value
     *
     * @return boolean
     */
    public function hasJobType(){
      return $this->_has(4);
    }
    
    /**
     * Clear <job_type> value
     *
     * @return \dtypes\LongRunningResponse
     */
    public function clearJobType(){
      return $this->_clear(4);
    }
    
    /**
     * Get <job_type> value
     *
     * @return string
     */
    public function getJobType(){
      return $this->_get(4);
    }
    
    /**
     * Set <job_type> value
     *
     * @param string $value
     * @return \dtypes\LongRunningResponse
     */
    public function setJobType( $value){
      return $this->_set(4, $value);
    }
  }
}

namespace dtypes {

  class VoidRequest extends \DrSlump\Protobuf\Message {


    /** @var \Closure[] */
    protected static $__extensions = array();

    public static function descriptor()
    {
      $descriptor = new \DrSlump\Protobuf\Descriptor(__CLASS__, 'dtypes.VoidRequest');

      foreach (self::$__extensions as $cb) {
        $descriptor->addField($cb(), true);
      }

      return $descriptor;
    }
  }
}

namespace dtypes {

  class VoidResponse extends \DrSlump\Protobuf\Message {

    /**  @var \dtypes\Status */
    public $status = null;
    

    /** @var \Closure[] */
    protected static $__extensions = array();

    public static function descriptor()
    {
      $descriptor = new \DrSlump\Protobuf\Descriptor(__CLASS__, 'dtypes.VoidResponse');

      // OPTIONAL MESSAGE status = 1
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 1;
      $f->name      = "status";
      $f->type      = \DrSlump\Protobuf::TYPE_MESSAGE;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $f->reference = '\dtypes\Status';
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
     * @return \dtypes\VoidResponse
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
     * @return \dtypes\VoidResponse
     */
    public function setStatus(\dtypes\Status $value){
      return $this->_set(1, $value);
    }
  }
}

