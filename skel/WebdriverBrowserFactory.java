import org.openqa.selenium.WebDriver;
import org.openqa.selenium.firefox.FirefoxDriver;
import org.openqa.selenium.chrome.ChromeDriver;
import org.openqa.selenium.firefox.FirefoxDriver;
import org.openqa.selenium.htmlunit.HtmlUnitDriver;
import org.openqa.selenium.remote.DesiredCapabilities;
import org.openqa.selenium.safari.SafariDriver;


import java.lang.Exception;
import java.util.HashMap;

public class WebdriverBrowserFactory {
    static WebDriver driver;

    public static void init()throws Exception {
        String webDriverBrowser = System.getenv("webdriver_driver");
        if(webDriverBrowser == null || webDriverBrowser.equalsIgnoreCase(""))
            webDriverBrowser = "firefox";
        
        HashMap<String, WebDriver> browsers = new HashMap<String, WebDriver>();
        browsers.put("firefox",new FirefoxDriver());
      
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
