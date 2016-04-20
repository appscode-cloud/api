<?php
include 'vendor/autoload.php';
include "../health.php";
include "../types.php";

$host = 'localhost:50051';
$client = new health\HealthClient($host,  [
  'credentials' => Grpc\ChannelCredentials::createInsecure(),
]);

$req = new dtypes\VoidRequest();
$call = $client->Status($req);
$call->getPeer();
list($resp, $status) = $call->wait();
echo json_encode($resp->getVersion());

