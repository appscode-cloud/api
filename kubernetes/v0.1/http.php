<?php
// DO NOT EDIT! Generated by Protobuf-PHP protoc plugin 1.0
// Source: google/api/http.proto
//   Date: 2016-04-19 16:27:31

namespace google\api {

  class HttpRule extends \DrSlump\Protobuf\Message {

    /**  @var string */
    public $get = null;
    
    /**  @var string */
    public $put = null;
    
    /**  @var string */
    public $post = null;
    
    /**  @var string */
    public $delete = null;
    
    /**  @var string */
    public $patch = null;
    
    /**  @var \google\api\CustomHttpPattern */
    public $custom = null;
    
    /**  @var string */
    public $body = null;
    
    /**  @var \google\api\HttpRule[]  */
    public $additional_bindings = array();
    

    /** @var \Closure[] */
    protected static $__extensions = array();

    public static function descriptor()
    {
      $descriptor = new \DrSlump\Protobuf\Descriptor(__CLASS__, 'google.api.HttpRule');

      // OPTIONAL STRING get = 2
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 2;
      $f->name      = "get";
      $f->type      = \DrSlump\Protobuf::TYPE_STRING;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $descriptor->addField($f);

      // OPTIONAL STRING put = 3
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 3;
      $f->name      = "put";
      $f->type      = \DrSlump\Protobuf::TYPE_STRING;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $descriptor->addField($f);

      // OPTIONAL STRING post = 4
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 4;
      $f->name      = "post";
      $f->type      = \DrSlump\Protobuf::TYPE_STRING;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $descriptor->addField($f);

      // OPTIONAL STRING delete = 5
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 5;
      $f->name      = "delete";
      $f->type      = \DrSlump\Protobuf::TYPE_STRING;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $descriptor->addField($f);

      // OPTIONAL STRING patch = 6
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 6;
      $f->name      = "patch";
      $f->type      = \DrSlump\Protobuf::TYPE_STRING;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $descriptor->addField($f);

      // OPTIONAL MESSAGE custom = 8
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 8;
      $f->name      = "custom";
      $f->type      = \DrSlump\Protobuf::TYPE_MESSAGE;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $f->reference = '\google\api\CustomHttpPattern';
      $descriptor->addField($f);

      // OPTIONAL STRING body = 7
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 7;
      $f->name      = "body";
      $f->type      = \DrSlump\Protobuf::TYPE_STRING;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $descriptor->addField($f);

      // REPEATED MESSAGE additional_bindings = 11
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 11;
      $f->name      = "additional_bindings";
      $f->type      = \DrSlump\Protobuf::TYPE_MESSAGE;
      $f->rule      = \DrSlump\Protobuf::RULE_REPEATED;
      $f->reference = '\google\api\HttpRule';
      $descriptor->addField($f);

      foreach (self::$__extensions as $cb) {
        $descriptor->addField($cb(), true);
      }

      return $descriptor;
    }

    /**
     * Check if <get> has a value
     *
     * @return boolean
     */
    public function hasGet(){
      return $this->_has(2);
    }
    
    /**
     * Clear <get> value
     *
     * @return \google\api\HttpRule
     */
    public function clearGet(){
      return $this->_clear(2);
    }
    
    /**
     * Get <get> value
     *
     * @return string
     */
    public function getGet(){
      return $this->_get(2);
    }
    
    /**
     * Set <get> value
     *
     * @param string $value
     * @return \google\api\HttpRule
     */
    public function setGet( $value){
      return $this->_set(2, $value);
    }
    
    /**
     * Check if <put> has a value
     *
     * @return boolean
     */
    public function hasPut(){
      return $this->_has(3);
    }
    
    /**
     * Clear <put> value
     *
     * @return \google\api\HttpRule
     */
    public function clearPut(){
      return $this->_clear(3);
    }
    
    /**
     * Get <put> value
     *
     * @return string
     */
    public function getPut(){
      return $this->_get(3);
    }
    
    /**
     * Set <put> value
     *
     * @param string $value
     * @return \google\api\HttpRule
     */
    public function setPut( $value){
      return $this->_set(3, $value);
    }
    
    /**
     * Check if <post> has a value
     *
     * @return boolean
     */
    public function hasPost(){
      return $this->_has(4);
    }
    
    /**
     * Clear <post> value
     *
     * @return \google\api\HttpRule
     */
    public function clearPost(){
      return $this->_clear(4);
    }
    
    /**
     * Get <post> value
     *
     * @return string
     */
    public function getPost(){
      return $this->_get(4);
    }
    
    /**
     * Set <post> value
     *
     * @param string $value
     * @return \google\api\HttpRule
     */
    public function setPost( $value){
      return $this->_set(4, $value);
    }
    
    /**
     * Check if <delete> has a value
     *
     * @return boolean
     */
    public function hasDelete(){
      return $this->_has(5);
    }
    
    /**
     * Clear <delete> value
     *
     * @return \google\api\HttpRule
     */
    public function clearDelete(){
      return $this->_clear(5);
    }
    
    /**
     * Get <delete> value
     *
     * @return string
     */
    public function getDelete(){
      return $this->_get(5);
    }
    
    /**
     * Set <delete> value
     *
     * @param string $value
     * @return \google\api\HttpRule
     */
    public function setDelete( $value){
      return $this->_set(5, $value);
    }
    
    /**
     * Check if <patch> has a value
     *
     * @return boolean
     */
    public function hasPatch(){
      return $this->_has(6);
    }
    
    /**
     * Clear <patch> value
     *
     * @return \google\api\HttpRule
     */
    public function clearPatch(){
      return $this->_clear(6);
    }
    
    /**
     * Get <patch> value
     *
     * @return string
     */
    public function getPatch(){
      return $this->_get(6);
    }
    
    /**
     * Set <patch> value
     *
     * @param string $value
     * @return \google\api\HttpRule
     */
    public function setPatch( $value){
      return $this->_set(6, $value);
    }
    
    /**
     * Check if <custom> has a value
     *
     * @return boolean
     */
    public function hasCustom(){
      return $this->_has(8);
    }
    
    /**
     * Clear <custom> value
     *
     * @return \google\api\HttpRule
     */
    public function clearCustom(){
      return $this->_clear(8);
    }
    
    /**
     * Get <custom> value
     *
     * @return \google\api\CustomHttpPattern
     */
    public function getCustom(){
      return $this->_get(8);
    }
    
    /**
     * Set <custom> value
     *
     * @param \google\api\CustomHttpPattern $value
     * @return \google\api\HttpRule
     */
    public function setCustom(\google\api\CustomHttpPattern $value){
      return $this->_set(8, $value);
    }
    
    /**
     * Check if <body> has a value
     *
     * @return boolean
     */
    public function hasBody(){
      return $this->_has(7);
    }
    
    /**
     * Clear <body> value
     *
     * @return \google\api\HttpRule
     */
    public function clearBody(){
      return $this->_clear(7);
    }
    
    /**
     * Get <body> value
     *
     * @return string
     */
    public function getBody(){
      return $this->_get(7);
    }
    
    /**
     * Set <body> value
     *
     * @param string $value
     * @return \google\api\HttpRule
     */
    public function setBody( $value){
      return $this->_set(7, $value);
    }
    
    /**
     * Check if <additional_bindings> has a value
     *
     * @return boolean
     */
    public function hasAdditionalBindings(){
      return $this->_has(11);
    }
    
    /**
     * Clear <additional_bindings> value
     *
     * @return \google\api\HttpRule
     */
    public function clearAdditionalBindings(){
      return $this->_clear(11);
    }
    
    /**
     * Get <additional_bindings> value
     *
     * @param int $idx
     * @return \google\api\HttpRule
     */
    public function getAdditionalBindings($idx = NULL){
      return $this->_get(11, $idx);
    }
    
    /**
     * Set <additional_bindings> value
     *
     * @param \google\api\HttpRule $value
     * @return \google\api\HttpRule
     */
    public function setAdditionalBindings(\google\api\HttpRule $value, $idx = NULL){
      return $this->_set(11, $value, $idx);
    }
    
    /**
     * Get all elements of <additional_bindings>
     *
     * @return \google\api\HttpRule[]
     */
    public function getAdditionalBindingsList(){
     return $this->_get(11);
    }
    
    /**
     * Add a new element to <additional_bindings>
     *
     * @param \google\api\HttpRule $value
     * @return \google\api\HttpRule
     */
    public function addAdditionalBindings(\google\api\HttpRule $value){
     return $this->_add(11, $value);
    }
  }
}

namespace google\api {

  class CustomHttpPattern extends \DrSlump\Protobuf\Message {

    /**  @var string */
    public $kind = null;
    
    /**  @var string */
    public $path = null;
    

    /** @var \Closure[] */
    protected static $__extensions = array();

    public static function descriptor()
    {
      $descriptor = new \DrSlump\Protobuf\Descriptor(__CLASS__, 'google.api.CustomHttpPattern');

      // OPTIONAL STRING kind = 1
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 1;
      $f->name      = "kind";
      $f->type      = \DrSlump\Protobuf::TYPE_STRING;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $descriptor->addField($f);

      // OPTIONAL STRING path = 2
      $f = new \DrSlump\Protobuf\Field();
      $f->number    = 2;
      $f->name      = "path";
      $f->type      = \DrSlump\Protobuf::TYPE_STRING;
      $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
      $descriptor->addField($f);

      foreach (self::$__extensions as $cb) {
        $descriptor->addField($cb(), true);
      }

      return $descriptor;
    }

    /**
     * Check if <kind> has a value
     *
     * @return boolean
     */
    public function hasKind(){
      return $this->_has(1);
    }
    
    /**
     * Clear <kind> value
     *
     * @return \google\api\CustomHttpPattern
     */
    public function clearKind(){
      return $this->_clear(1);
    }
    
    /**
     * Get <kind> value
     *
     * @return string
     */
    public function getKind(){
      return $this->_get(1);
    }
    
    /**
     * Set <kind> value
     *
     * @param string $value
     * @return \google\api\CustomHttpPattern
     */
    public function setKind( $value){
      return $this->_set(1, $value);
    }
    
    /**
     * Check if <path> has a value
     *
     * @return boolean
     */
    public function hasPath(){
      return $this->_has(2);
    }
    
    /**
     * Clear <path> value
     *
     * @return \google\api\CustomHttpPattern
     */
    public function clearPath(){
      return $this->_clear(2);
    }
    
    /**
     * Get <path> value
     *
     * @return string
     */
    public function getPath(){
      return $this->_get(2);
    }
    
    /**
     * Set <path> value
     *
     * @param string $value
     * @return \google\api\CustomHttpPattern
     */
    public function setPath( $value){
      return $this->_set(2, $value);
    }
  }
}

