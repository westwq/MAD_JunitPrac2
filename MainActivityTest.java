package sg.edu.np.mad.madpractical;

import android.app.Activity;
import android.content.Intent;
import android.util.Log;

import org.junit.After;
import org.junit.Before;
import org.junit.BeforeClass;
import org.junit.Rule;
import org.junit.Test;
import org.junit.runner.RunWith;

import androidx.test.core.app.ActivityScenario;
import androidx.test.ext.junit.rules.ActivityScenarioRule;
import androidx.test.ext.junit.runners.AndroidJUnit4;

import static androidx.test.espresso.Espresso.onView;
import static androidx.test.espresso.action.ViewActions.click;
import static androidx.test.espresso.assertion.ViewAssertions.matches;
import static androidx.test.espresso.matcher.ViewMatchers.isDisplayed;
import static androidx.test.espresso.matcher.ViewMatchers.withId;

@RunWith(AndroidJUnit4.class)

public class MainActivityTest {


    @Rule
    public ActivityScenarioRule<MainActivity> activityRule =
            new ActivityScenarioRule<>(MainActivity.class);
    @Before
    public void setUp() throws Exception {
        activityRule.getScenario();
    }

    @Test
    public void changeText_sameActivity() {
        // Type text and then press the button
        onView(withId(R.id.btnFollow)).perform(click());
        onView(withId(R.id.btnFollow)).perform(click());
        onView(withId(R.id.btnFollow)).perform(click());
        onView(withId(R.id.btnFollow)).perform(click());
        onView(withId(R.id.btnFollow)).perform(click());
        onView(withId(R.id.btnFollow)).check(matches(isDisplayed()));
    }
    
    @After
    public void tearDown() throws Exception {
    }
}
