<?php
// DO NOT EDIT! Generated by Protobuf-PHP protoc plugin 1.0
// Source: google/api/annotations.proto
//   Date: 2016-04-20 05:49:30

namespace {
  \google\protobuf\MethodOptions::extension(function(){
      // OPTIONAL MESSAGE google\api\http = 72295728
    $f = new \DrSlump\Protobuf\Field();
    $f->number    = 72295728;
    $f->name      = "google\api\http";
    $f->type      = \DrSlump\Protobuf::TYPE_MESSAGE;
    $f->rule      = \DrSlump\Protobuf::RULE_OPTIONAL;
    $f->reference = '\google\api\HttpRule';
    return $f;
  });
}