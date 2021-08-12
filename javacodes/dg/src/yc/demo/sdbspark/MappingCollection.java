package yc.demo.sdbspark;

public class MappingCollection {
    public static void link() {
        String mappingTable1 =
                "CREATE TABLE employee1 " +
                        "USING com.sequoiadb.spark " +
                        "OPTIONS( " +
                        "host 'sdbserver1:11810', " +
                        "collectionspace 'quiz', " +
                        "collection 'employee1', " +
                        "user 'sdbadmin'," +
                        "password 'sdbadmin'" +
                        ")";
        String mappingTable2 =
                "CREATE TABLE employee2 " +
                        "USING com.sequoiadb.spark " +
                        "OPTIONS( " +
                        "host 'sdbserver1:11810', " +
                        "collectionspace 'quiz', " +
                        "collection 'employee2', " +
                        "user 'sdbadmin'," +
                        "password 'sdbadmin'" +
                        ")";
        JDBCUtil.doDDL(mappingTable1);
        JDBCUtil.doDDL(mappingTable2);
    }
}
