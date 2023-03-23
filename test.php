<?php
function passCheck($pass)
{
    //echo $pass;
    $url = 'http://localhost:227/api/v1/hash';

    //The data you want to send via POST
    $fields = [
        'hash' => $pass,
    ];

    $result = httpPost($url, $fields);
    echo $result , "\n";
}

function httpPost($url, $data)
{
    $curl = curl_init($url);
    curl_setopt($curl, CURLOPT_POST, true);
    curl_setopt($curl, CURLOPT_POSTFIELDS, $data);
    curl_setopt($curl, CURLOPT_RETURNTRANSFER, true);
    $response = curl_exec($curl);
    curl_close($curl);
    return $response;
}

passCheck("password");