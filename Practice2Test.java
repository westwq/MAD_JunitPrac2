package sg.edu.np.mad.madpractical;

import org.junit.Test;

import static org.junit.Assert.*;

public class UserTest{
    private User user;

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
