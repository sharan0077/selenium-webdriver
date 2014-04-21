import com.thoughtworks.twist2.AfterSpec;
import com.thoughtworks.twist2.BeforeSpec;


public class WebdriverHooks {

    @BeforeSpec
    public void initializeBrowser()throws Exception{
        WebdriverBrowserFactory.init();
    }

    @AfterSpec
    public void killBrowser(){
        WebdriverBrowserFactory.kill();
    }

}
