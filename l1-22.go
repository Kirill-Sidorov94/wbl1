package main

import (
    "fmt"
    "math/big"
)

type bigNumsCalculator struct{
	operation map[string]bool
}

func newBigNumsCalculator() *bigNumsCalculator {
    return &bigNumsCalculator{
    	operation: map[string]bool{
	        "+": true,
	        "-": true,
	        "*": true,
	        "/": true,
	    },
    }
}

func (c *bigNumsCalculator) calculate(a any, op string, b any) (any, error) {
   	if _, ok := c.operation[op]; !ok {
   		return nil, fmt.Errorf("bigNumsCalculator.calculate(): invalid operation: %s", op)
   	} 

    aBig, err := c.toBigNumber(a)
    if err != nil {
        return nil, fmt.Errorf("bigNumsCalculator.calculate(): left num: %w", err)
    }

    bBig, err := c.toBigNumber(b)
    if err != nil {
        return nil, fmt.Errorf("bigNumsCalculator.calculate(): right num: %w", err)
    }

    result, err := c.calculateBig(aBig, bBig, op)
    if err != nil {
    	return nil, fmt.Errorf("bigNumsCalculator.calculate(): %w", err)
    }

    return result, nil
}

func (c *bigNumsCalculator) toBigNumber(value any) (any, error) {
    switch v := value.(type) {
    case int64:
        return big.NewInt(v), nil
    case float64:
        return big.NewFloat(v), nil
    case string:
        return c.parseStringToBigNumber(v)
    case *big.Int, *big.Float:
    	return v, nil
    default:
        return nil, fmt.Errorf("bigNumsCalculator.toBigNumber(): invalid type %T, expected int64, float64 or string", value)
    }
}

func (c *bigNumsCalculator) parseStringToBigNumber(s string) (any, error) {
    if bigInt, ok := new(big.Int).SetString(s, 10); ok {
        return bigInt, nil
    }
    
    if bigFloat, ok := new(big.Float).SetString(s); ok {
        return bigFloat, nil
    }
    
    return nil, fmt.Errorf("bigNumsCalculator.parseStringToBigNumber(): cannot parse string as number: %s", s)
}

func (c *bigNumsCalculator) calculateBig(a, b any, op string) (any, error) {
    switch aVal := a.(type) {
    case *big.Int:
        if bVal, ok := b.(*big.Int); ok {
            return c.calculateBigInt(aVal, bVal, op)
        }
        return nil, fmt.Errorf("bigNumsCalculator.calculateBig(): type mismatch: big.Int and %T", b)
    
    case *big.Float:
        if bVal, ok := b.(*big.Float); ok {
            return c.calculateBigFloat(aVal, bVal, op)
        }
        return nil, fmt.Errorf("bigNumsCalculator.calculateBig(): type mismatch: big.Float and %T", b)
    
    default:
        return nil, fmt.Errorf("bigNumsCalculator.calculateBig(): unsupported type: %T", a)
    }
}

func (c *bigNumsCalculator) calculateBigInt(a, b *big.Int, op string) (*big.Int, error) {
	result := new(big.Int)

    switch op {
    case "+": 
    	return result.Add(a, b), nil
    case "-": 
    	return result.Sub(a, b), nil  
    case "*": 
    	return result.Mul(a, b), nil
    case "/": 
        if b.Sign() == 0 { 
        	return nil, fmt.Errorf("bigNumsCalculator.calculateBigInt(): division by zero") 
        }

        return result.Div(a, b), nil
    default: 
    	return nil, fmt.Errorf("bigNumsCalculator.calculateBigInt(): unknown operation")
    }
}

func (c *bigNumsCalculator) calculateBigFloat(a, b *big.Float, op string) (*big.Float, error) {
	result := new(big.Float)

    switch op {
    case "+": 
    	return result.Add(a, b), nil
    case "-": 
    	return result.Sub(a, b), nil
    case "*": 
    	return result.Mul(a, b), nil  
    case "/":
        if b.Sign() == 0 { 
        	return nil, fmt.Errorf("bigNumsCalculator.calculateBigFloat(): division by zero") 
        }

        return result.Quo(a, b), nil
    default: 
    	return nil, fmt.Errorf("bigNumsCalculator.calculateBigFloat(): unknown operation")
    }
}

func workingWithBigNums() {
	calc := newBigNumsCalculator()

	bigIntNum1, _ := new(big.Int).SetString("283462847362817263548172635481726354", 10)
	bigIntNum2, _ := new(big.Int).SetString("115792089237316195423570985008687907853269984665640564039457584007913129639935", 10)
	bigIntAddResult, _ := calc.calculate(bigIntNum1, "+", bigIntNum2)
	fmt.Println(bigIntAddResult)

	bigFloatNum1, _ := new(big.Float).SetString("283462847362817263548172635481726354.12345678901234567890")
	bigFloatNum2, _ := new(big.Float).SetString("115792089237316195423570985008687907853269984665640564039457584007913129639935.98765432109876543210")
	bigFloatSubResult, _ := calc.calculate(bigFloatNum1, "-", bigFloatNum2)
	fmt.Println(bigFloatSubResult)

	bigIntStr1 := "283462847362817263548172635481726354"
	bigIntStr2 := "115792089237316195423570985008687907853269984665640564039457584007913129639935"
	bigFloatMulResult, _ := calc.calculate(bigIntStr1, "*", bigIntStr2)
	fmt.Println(bigFloatMulResult)

	int64Num1 := int64(9223372036854775807)
	int64Num2 := int64(-9223372036854775808)
	int64DivResult, _ := calc.calculate(int64Num1, "/", int64Num2)
	fmt.Println(int64DivResult)

	float64Num1 := 1.7976931348623157e308
	float64Num2 := 2.2250738585072014e-308
	float64QuoResult, _ := calc.calculate(float64Num1, "/", float64Num2)
	fmt.Println(float64QuoResult)

	str1 := "drop datebase users"
	bigIntAddResult, err := calc.calculate(bigIntNum1, "+", str1)
	if err != nil {
		fmt.Println(err.Error())
	}
}