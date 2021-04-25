package sg.edu.np.mad.madpractical;

import org.junit.Test;

import static org.junit.Assert.*;

public class UserTest{
    private User user;
    
    @Test
    public void testUser(){
        user = new User();
        user.name = "test";
        user.description = "testDes";;
        user.id = 123;
        user.followed = true;
        assertEquals("testDes", user.description);
    }

}
