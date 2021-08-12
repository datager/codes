package gaussDB;

import java.sql.*;

public class CopyManagerTest {
    static Connection conn;
    static Statement stmt;

    public static void main(String[] args) {
        try {
            Class.forName("org.postgresql.Driver");
            conn = DriverManager.getConnection("jdbc:postgresql://192.168.2.142:5432/statistics", "postgres", "Zstvgcs@9102");
            stmt = conn.createStatement();
            String sql = "SELECT * from vehicle_capture_hour_count limit 10";
            ResultSet rs = stmt.executeQuery(sql);

            while (rs.next()) {
                Long ts = rs.getLong("ts");
                Long oc = rs.getLong("org_code");
                Integer sensorIntID = rs.getInt("sensor_int_id");
                String sensorID = rs.getString("sensor_id");
                Long cnt = rs.getLong("count");

                System.out.printf("%d-%d-%d-%s-%d", ts, oc, sensorIntID, sensorID, cnt);
            }
            rs.close();
            stmt.close();
            conn.close();
        } catch (ClassNotFoundException e) {
            e.printStackTrace();
        } catch (SQLException throwables) {
            throwables.printStackTrace();
        } finally {
            try {
                if (stmt != null) {
                    stmt.close();
                }
            } catch (SQLException ignored) {

            }

            try {
                if (conn != null) {
                    conn.close();
                }
            } catch (SQLException se) {
                se.printStackTrace();
            }
            System.out.println("Goodbye!");
        }
    }
}

