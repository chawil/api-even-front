<?php
if (!isset($_GET['input']) || $_GET['input'] == "" ) {
    include 'form.html';
} else {
    echo $_GET['input'];
}
?>

