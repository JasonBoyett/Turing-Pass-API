package encrypt

import ("testing")

func TestVariance(t *testing.T){
  enitialResult, err := Encrypt("foo", "bar", true, 16)
  if err != nil {
    t.Error("Execution Error: ", err)
  }
  secondResult, err := Encrypt("bar", "foo", true, 16)
  if err != nil {
    t.Error("Execution Error: ", err)
  }
  if enitialResult == secondResult {
    t.Errorf("FAIL Variance Failure\ninitial %s\nterms reversed: %s\n", enitialResult, secondResult)
  }
  if enitialResult != secondResult {
    t.Logf("PASS Variance Success\ninitial: %s\nterms reversed: %s\n", enitialResult, secondResult)
  }
}

func TestVarianceLong(t *testing.T){
  enitialResult, err := Encrypt("foo", "bar", true, 100)
  if err != nil {
    t.Error("Execution Error: ", err)
  }
  secondResult, err := Encrypt("bar", "foo", true, 100)
  if err != nil {
    t.Error("Execution Error: ", err)
  }
  if enitialResult == secondResult {
    t.Errorf("FAIL Variance Failure\ninitial %s\nterms reversed: %s\n", enitialResult, secondResult)
  }
  if enitialResult != secondResult {
    t.Logf("PASS Variance Success\ninitial: %s\nterms reversed: %s\n", enitialResult, secondResult)
  }
}

//I have run this test with 1_000_000_000 iterations and it has never failed
//I'm decreasing that number to 1_000_000 for the sake of time in the future
func TestConsistancy(t *testing.T){
  enitialResult, err := Encrypt("foo", "bar", true, 16)
  if err != nil {
    t.Error("Execution Error: ", err)
  }

  for i := 0; i < 1_000_000; i++ {
    test, err := Encrypt("foo", "bar", true, 16)
    if err != nil {
      t.Error("Execution Error: ", err)
    }
    if test != enitialResult {
      t.Errorf("FAIL Consistancy Failure\ninitial %s\niteration %s\n", enitialResult, test)
      return
    }
  }
  t.Logf("PASS Consistancy Success")
}

func TestConsistancyNoSymbols(t *testing.T){
  enitialResult, err := Encrypt("foo", "bar", false, 16)
  if err != nil {
    t.Error("Execution Error: ", err)
  }

  for i := 0; i < 1_000_000; i++ {
    test, err := Encrypt("foo", "bar", false, 16)
    if err != nil {
      t.Error("Execution Error: ", err)
    }
    if test != enitialResult {
      t.Errorf("FAIL Consistancy Failure\ninitial %s\niteration %s\n", enitialResult, test)
      return
    }
  }
  t.Logf("PASS Consistancy Success")
}
