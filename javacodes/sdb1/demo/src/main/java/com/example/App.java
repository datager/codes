package com.example;

import java.sql.*;

public class App {
static {
    try {
        Class.forName("com.mysql.jdbc.Driver");
    } catch (ClassNotFoundException e) {
        e.printStackTrace();
    }
}

public static void main(String[] args) throws SQLException {
    String hostName = "192.168.2.48";
    String port = "3306";
    String databaseName = "db";
    String myUser = "root";
    String myPasswd = "123456";
    String url = "jdbc:mysql://" + hostName + ":" + port + "/" + databaseName + "?useSSL=false";
    Connection conn = DriverManager.getConnection(url, myUser, myPasswd);

    System.out.println("---INSERT---");
    String sql = "INSERT INTO tb VALUES(?,?,?)";
    PreparedStatement ins = conn.prepareStatement(sql);
    ins.setInt(1, 1);
    ins.setString(2, "Peter");
    ins.setString(3, "Parcker");
    ins.executeUpdate();

    System.out.println("---UPDATE---");
    sql = "UPDATE tb SET first_name=? WHERE id = ?";
    PreparedStatement upd = conn.prepareStatement(sql);
    upd.setString(1, "Stephen");
    upd.setInt(2, 1);
    upd.executeUpdate();

    System.out.println("---SELECT---");
    Statement stmt = conn.createStatement();
    sql = "SELECT * FROM tb";
    ResultSet rs = stmt.executeQuery(sql);
    boolean isHeaderPrint = false;
    while (rs.next()) {
        ResultSetMetaData md = rs.getMetaData();
        int col_num = md.getColumnCount();
        if (!isHeaderPrint) {
            System.out.print("|");
            for (int i = 1; i <= col_num; i++) {
                System.out.print(md.getColumnName(i) + "\t|");
                isHeaderPrint = true;
            }
        }
        System.out.println();
        System.out.print("|");
        for (int i = 1; i <= col_num; i++) {
            System.out.print(rs.getString(i) + "\t|");
        }
        System.out.println();
    }
    stmt.close();

    System.out.println("---DELETE---");
    sql = "DELETE FROM tb WHERE id = ?";
    PreparedStatement del = conn.prepareStatement(sql);
    del.setInt(1, 1);
    del.executeUpdate();

    conn.close();
  }
}