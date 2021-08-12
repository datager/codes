package yc.demo.sdbspark;

public class CountEmployee {
    public static void CountBySex() {
        String aggregate =
                "CREATE TABLE employee_count USING com.sequoiadb.spark  " +
                        "OPTIONS( " +
                        "host 'sdbserver1:11810', " +
                        "collectionspace 'quiz', " +
                        "collection 'employee_count' " +
                        ")AS( " +
                        "SELECT sex," +
                        "num" +
                        "FROM (" +
                        "        SELECT sex," +
                        "        count(*) AS num" +
                        "FROM employee1 e" +
                        "GROUP BY sex" +
                        ") AS avg_table);";
        JDBCUtil.doDML(aggregate);
    }
}
