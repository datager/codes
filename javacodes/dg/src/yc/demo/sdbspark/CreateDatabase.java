package yc.demo.sdbspark;

import java.sql.*;

import static yc.demo.sdbspark.JDBCUtil.releaseSource;

public class CreateDatabase {
    private static final String url = "jdbc:hive2//sdbserver1:10000/default";
    private static final String username = "sdbadmin";
    private static final String password = "";

    public static void createDatabase() {
        try {
            Connection connection = DriverManager.getConnection(url, username, password);
            String createDatabaseSQL = "CREATE DATABASE IF NOT EXISTS quiz";
            Statement statement = connection.createStatement();
            statement.execute(createDatabaseSQL);
            String showDatabases = "SHOW DATABASES";
            ResultSet resultSet = statement.executeQuery(showDatabases);
//            ResultFormat.printResultSet(resultSet);
        } catch (SQLException e) {
            e.printStackTrace();
        }
//        String databaseName = "quiz";
//        // Get Hive JDBC connection
//        Connection connection = null;
//        try {
//            try {
//                Class.forName("org.apache.hive.jdbc.HiveDriver");
//                connection = DriverManager.getConnection(url, username, password);
//            } catch (ClassNotFoundException e) {
//                e.printStackTrace();
//            }
//        } catch (Exception throwables) {
//            throwables.printStackTrace();
//        }
//        // Initialize Statement
//        Statement statement = null;
//        // Initialize ResultSet
//        ResultSet resultSet = null;
//        try {
//            // SQL statement of create database
//            String createDatabaseSQL = "CREATE DATABASE IF NOT EXISTS " + databaseName;
//            // Create Statement
//            statement = connection.createStatement();
//            // Execute the SQL statement of create database
//            statement.execute(createDatabaseSQL);
//            // View the database statement
//            String showDatabases = "SHOW DATABASES";
//            // Execute the view database statement to get the result set
//            resultSet = statement.executeQuery(showDatabases);
//            // Call the predefined result set to beautify the utility class and print the result set
//            ResultFormat.printResultSet(resultSet);
//        } catch (SQLException e) {
//            e.printStackTrace();
//        } finally {
//            // Release Statement and Connection source
//            releaseSource(resultSet, statement, connection);
//        }
    }
}