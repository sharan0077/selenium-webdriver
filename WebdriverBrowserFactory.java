import org.openqa.selenium.WebDriver;
import org.openqa.selenium.firefox.FirefoxDriver;
import org.openqa.selenium.remote.DesiredCapabilities;

public class WebdriverBrowserFactory {
    static FirefoxDriver driver;

    public static void init(DesiredCapabilities capabilities) {
        driver = new FirefoxDriver(capabilities);
    }

    public static WebDriver getBroser(){
        return driver;
    }

    public static void kill(){
        driver.kill();
    }

}
