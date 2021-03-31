mysql code
user_id INT AUTO_INCREMENT,
username varchar(15),
gender varchar(10),
email varchar(255),
PRIMARY KEY(user_id)

nano /var/www/web/user.php

<?php
$user = "chen";
$password = "12345678";
$database = "web";
$table = "user";

$db = new PDO("mysql:host=localhost;dbname=$database", $user, $password);
echo "<h2>Web Service</h2><ol>";

foreach($db->query("SELECT * FROM $table") as $row) {
  if($_POST["id"] == $row['user_id']) {
    echo "<li>" . "Username: " . $row['username'] . ", Gender: " . $row['gender>
  }
}
echo "</ol>";
?>
<p><a href="/user.html">back</a></p>


nano /var/www/web/user.html
<html>
<body>

<form action="user.php" method="post">
User id: <input type="text" name="id"><br>
<input type="submit">
</form>

</body>
</html>


http://10.0.2.15/user.html

