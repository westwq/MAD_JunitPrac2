import org.junit.Test;

public class UserTest{
    private User user;

    @BeforeEach
    public void setUp() throws Exception{
        user = new User();
        user.name = "test";
        user.description = "testDes";;
        user.id = 123;
        user.followed = true;
    }

    @Test
    public void testUser(){
        assertEquals("testDes", user.description);
    }

}
