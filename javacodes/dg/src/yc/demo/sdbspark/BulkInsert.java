package yc.demo.sdbspark;

public class BulkInsert {
    public static void insert() {
        JDBCUtil.doDML("INSERT INTO employee1 SELECT * FROM employee2");
    }
}
