import org.openqa.selenium.WebDriver;
import org.openqa.selenium.chrome.ChromeDriver;
import org.openqa.selenium.firefox.FirefoxDriver;
import org.openqa.selenium.htmlunit.HtmlUnitDriver;
import org.openqa.selenium.remote.DesiredCapabilities;
import org.openqa.selenium.safari.SafariDriver;

import java.lang.Exception;
import java.util.HashMap;

public class WebdriverBrowserFactory {
    static WebDriver driver;
    public static void init(DesiredCapabilities capabilities)throws Exception{

        String webDriverBrowser = System.getenv("webdriver_browser");

        if(webDriverBrowser == null)
            webDriverBrowser = "firefox";
        
        HashMap<String, WebDriver> browsers = new HashMap<String, WebDriver>();
        browsers.put("chrome",new ChromeDriver(capabilities));
        browsers.put("firefox",new FirefoxDriver(capabilities));
        browsers.put("internetexplorer",new ChromeDriver(capabilities));
        browsers.put("htmlunitdriver",new HtmlUnitDriver(capabilities));
        browsers.put("safaridriver",new SafariDriver(capabilities));

        driver =   browsers.get(webDriverBrowser.toLowerCase());
        if (driver == null) {
            throw new RuntimeException(webDriverBrowser + " is invalid");
        }
    }

    public static WebDriver getBroser(){
        return driver;
    }

    public static void kill(){
//        driver.kill();
    }

}
