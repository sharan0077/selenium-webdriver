import com.thoughtworks.twist2.AfterSuite;
import com.thoughtworks.twist2.BeforeSuite;


public class WebdriverHooks {

    @BeforeSuite
    public void initializeBrowser()throws Exception{
        WebdriverBrowserFactory.init();
    }

    @AfterSuite
    public void killBrowser(){
        WebdriverBrowserFactory.kill();
    }

}
