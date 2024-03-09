import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.assertEquals;

public class customerDataTest { 
    @Test
    public void testSETDATA() { 
        customerData data = new customerData(); 
        data.SETDATA(1, "John Doe"); 
        assertEquals("John Doe", data.getCustomerName());
    }
}