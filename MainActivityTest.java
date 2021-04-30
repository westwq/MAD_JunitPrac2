package sg.edu.np.mad.madpractical;

import androidx.test.espresso.ViewAction;
import androidx.test.espresso.matcher.ViewMatchers;
import androidx.test.ext.junit.rules.ActivityScenarioRule;

import org.junit.After;
import org.junit.Before;
import org.junit.Rule;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.junit.runners.JUnit4;

import static androidx.test.espresso.Espresso.onView;
import static androidx.test.espresso.action.ViewActions.click;
import static org.junit.Assert.*;
@RunWith(JUnit4.class)
public class MainActivityTest {
    @Rule
    public ActivityScenarioRule<MainActivity> activityRule = new ActivityScenarioRule<>(MainActivity.class);


    @Before
    public void setUp() throws Exception {
    }

    @Test
    public void pressImageView(){
        onView(ViewMatchers.withId(R.id.btnFollow)).perform(click());
        onView(ViewMatchers.withId(R.id.btnFollow)).perform(click());


    }

    @After
    public void tearDown() throws Exception {
    }
}
