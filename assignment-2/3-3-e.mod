set BEGIN; # plant
set ORIG; # mill
set DEST; # destinations
param mat_supply {BEGIN} >= 0; # amount of material at plants
param supply {ORIG} >= 0; # amounts available at origins
param demand {DEST} >= 0; # amounts required at destinations
param supply_pct >= 0, <= 100;
param demand_pct >= 0, <= 100;
check: sum {h in BEGIN} mat_supply[h] >= sum {i in ORIG} supply[i] <= sum {j in DEST} demand[j];
param mat_cost {BEGIN,ORIG} >= 0; # shipment costs per unit from plant
param cost {ORIG,DEST} >= 0; # shipment costs per unit to dest
var Mat_Trans {BEGIN,ORIG} >= 0; # amount for the plant to ship to factory
var Trans {ORIG,DEST} >= 0; # units to be shipped
param scrap_pct {ORIG} >= 0, <= 100;
param scrap_val >= 0;
minimize Total_Cost:
sum {h in BEGIN, i in ORIG} mat_cost[h, i] * Mat_Trans[h, i] +  sum {i in ORIG, j in DEST} cost[i,j] * Trans[i,j] - sum {i in ORIG} (supply[i] * scrap_pct[i]/100 * scrap_val);

subject to MaterialsSupply {h in BEGIN}:
sum {i in ORIG} Mat_Trans[h,i] == mat_supply[h];
subject to MaterialsDemand {i in ORIG}:
sum {h in BEGIN} Mat_Trans[h,i] == supply[i] * (100 - scrap_pct)/100;

subject to Supply {i in ORIG}:
sum {j in DEST} Trans[i,j] == supply[i] * (100 - scrap_pct)/100;
subject to Demand {j in DEST}:
sum {i in ORIG} Trans[i,j] == demand[j];

subject to SupplyPct {h in BEGIN, i in ORIG}:
Mat_Trans[h,i] <= (supply_pct/100) * mat_supply[h];
subject to DemandPct {h in BEGIN, i in ORIG}:
Mat_Trans[h,i] <= (demand_pct/100) * supply[i] * (100 - scrap_pct)/100;
