# trading-system (WIP)

## Database(mysql):
* name: trading_service
  * table: trading_order
    * columns:
      * ID (int, primary Key)
      * owner (string)
      * type (int, 1: Buy, 2:Sell)
      * quantity (int)
      * price (int)
      * price_policy (int, 1: Limit Price, 2: Market Price)


* Reference:
  1. https://www.cmegroup.com/confluence/display/EPICSANDBOX/Supported+Matching+Algorithms
  2. https://tianpan.co/notes/161-designing-stock-exchange
  3. https://github.com/go-programming-tour-book/blog-service
