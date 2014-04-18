import com.thoughtworks.twist2.AfterSpec;
import com.thoughtworks.twist2.BeforeSpec;
import org.openqa.selenium.remote.DesiredCapabilities;

public class WebdriverHooks {

    @BeforeSpec
    public void initializeBrowser(){
        WebdriverBrowserFactory.init(DesiredCapabilities.firefox());
    }

    @AfterSpec
    public void killBrowser(){
        WebdriverBrowserFactory.kill();
    }

}
