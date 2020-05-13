package emical;

import javax.jws.WebMethod;
import javax.jws.WebParam;
import javax.jws.WebService;

@WebService
public class Myclass {
    @WebMethod
    public float emi_calculator(@WebParam(name = "arg0") float p, @WebParam(name = "arg1") float r,
                                @WebParam(name = "arg2") float t) 
            { 
                    float emi; 
            
                    r = r / (12 * 100); // one month interest 
                    t = t * 12; // one month period 
                    emi = (p * r * (float)Math.pow(1 + r, t)) 
                                    / (float)(Math.pow(1 + r, t) - 1); 
            
                    return emi;
    }
}
