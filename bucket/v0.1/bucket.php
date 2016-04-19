<?php
// DO NOT EDIT! Generated by Protobuf-PHP protoc plugin 1.0
// Source: bucket.proto
//   Date: 2016-04-19 16:27:19

namespace bucket {

  class BucketListRequest extends \DrSlump\Protobuf\Message {

    /**  @var string */
    public $cloud_credential = null;
    

    /** @var \Closure[] */
    protected static $__extensions = array();

    public static function descriptor()
    {
      $descriptor = new \DrSlump\Protobuf\Descriptor(__CLASS__, 'bucket.BucketListRequest');

      // OPTIONAL STRING cloud_credential = 1
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 1;
      $f->name      = "cloud_credential";
      $f->type      = \DrSlump\Protobuf::TYPE_STRING;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $descriptor->addField($f);

      foreach (self::$__extensions as $cb) {
        $descriptor->addField($cb(), true);
      }

      return $descriptor;
    }

    /**
     * Check if <cloud_credential> has a value
     *
     * @return boolean
     */
    public function hasCloudCredential(){
      return $this->_has(1);
    }
    
    /**
     * Clear <cloud_credential> value
     *
     * @return \bucket\BucketListRequest
     */
    public function clearCloudCredential(){
      return $this->_clear(1);
    }
    
    /**
     * Get <cloud_credential> value
     *
     * @return string
     */
    public function getCloudCredential(){
      return $this->_get(1);
    }
    
    /**
     * Set <cloud_credential> value
     *
     * @param string $value
     * @return \bucket\BucketListRequest
     */
    public function setCloudCredential( $value){
      return $this->_set(1, $value);
    }
  }
}

namespace bucket {

  class BucketListResponse extends \DrSlump\Protobuf\Message {

    /**  @var string[]  */
    public $name = array();
    
    /**  @var \dtypes\Status */
    public $status = null;
    

    /** @var \Closure[] */
    protected static $__extensions = array();

    public static function descriptor()
    {
      $descriptor = new \DrSlump\Protobuf\Descriptor(__CLASS__, 'bucket.BucketListResponse');

      // REPEATED STRING name = 1
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 1;
      $f->name      = "name";
      $f->type      = \DrSlump\Protobuf::TYPE_STRING;
      $f->rule      = \DrSlump\Protobuf::RULE_REPEATED;
      $descriptor->addField($f);

      // OPTIONAL MESSAGE status = 2
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 2;
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
     * Check if <name> has a value
     *
     * @return boolean
     */
    public function hasName(){
      return $this->_has(1);
    }
    
    /**
     * Clear <name> value
     *
     * @return \bucket\BucketListResponse
     */
    public function clearName(){
      return $this->_clear(1);
    }
    
    /**
     * Get <name> value
     *
     * @param int $idx
     * @return string
     */
    public function getName($idx = NULL){
      return $this->_get(1, $idx);
    }
    
    /**
     * Set <name> value
     *
     * @param string $value
     * @return \bucket\BucketListResponse
     */
    public function setName( $value, $idx = NULL){
      return $this->_set(1, $value, $idx);
    }
    
    /**
     * Get all elements of <name>
     *
     * @return string[]
     */
    public function getNameList(){
     return $this->_get(1);
    }
    
    /**
     * Add a new element to <name>
     *
     * @param string $value
     * @return \bucket\BucketListResponse
     */
    public function addName( $value){
     return $this->_add(1, $value);
    }
    
    /**
     * Check if <status> has a value
     *
     * @return boolean
     */
    public function hasStatus(){
      return $this->_has(2);
    }
    
    /**
     * Clear <status> value
     *
     * @return \bucket\BucketListResponse
     */
    public function clearStatus(){
      return $this->_clear(2);
    }
    
    /**
     * Get <status> value
     *
     * @return \dtypes\Status
     */
    public function getStatus(){
      return $this->_get(2);
    }
    
    /**
     * Set <status> value
     *
     * @param \dtypes\Status $value
     * @return \bucket\BucketListResponse
     */
    public function setStatus(\dtypes\Status $value){
      return $this->_set(2, $value);
    }
  }
}

namespace bucket {

  class BucketsClient extends \Grpc\BaseStub {

    public function __construct($hostname, $opts) {
      parent::__construct($hostname, $opts);
    }
    /**
     * @param bucket\BucketListRequest $input
     */
    public function List(\bucket\BucketListRequest $argument, $metadata = array(), $options = array()) {
      return $this->_simpleRequest('/bucket.Buckets/List', $argument, '\bucket\BucketListResponse::deserialize', $metadata, $options);
    }
  }
}
