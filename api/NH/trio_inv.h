/************************************************************************************
	����

	�� �ڷ�� 2013�� 10�� 15�� ���� �ڷ��̸� ���� ����� ���ɼ��� �ֽ��ϴ�.
	�ڷ� ������ ���� ���� ��� ����ü�� ������� �ʾҴ��� Ȯ���Ͻñ� �ٶ��ϴ�.

	�ֽ� �ڷ�� ���������� ���� �ȳ��Ǹ� �ڵ� �ȳ�(OpenAPI Login��)�� �ϰ� ������ 
	�Խø� �� Ȯ���Ͻñ� �ٶ��ϴ�.

************************************************************************************/

typedef struct tagc1101InBlock    //�⺻�Է�
{
	char formlang                         [  1];	char _formlang;                           //�ѿ�����             
	char code                             [  6];	char _code;                               //�����ڵ�             
} Tc1101InBlock;

typedef struct tagc1101OutBlock    //���񸶽�Ÿ�⺻�ڷ�
{
	char code                             [  6];	char _code;                               //�����ڵ�             
	char hname                            [ 13];	char _hname;                              //�����               
	char price                            [  7];	char _price;                              //���簡               
	char sign                             [  1];	char _sign;                               //�����ȣ             
	char change                           [  6];	char _change;                             //�����               
	char chrate                           [  5];	char _chrate;                             //�����               
	char offer                            [  7];	char _offer;                              //�ŵ�ȣ��             
	char bid                              [  7];	char _bid;                                //�ż�ȣ��             
	char volume                           [  9];	char _volume;                             //�ŷ���               
	char volrate                          [  6];	char _volrate;                            //�ŷ�����             
	char yurate                           [  5];	char _yurate;                             //������ȸ����         
	char value                            [  9];	char _value;                              //�ŷ����             
	char uplmtprice                       [  7];	char _uplmtprice;                         //���Ѱ�               
	char high                             [  7];	char _high;                               //���߰�             
	char open                             [  7];	char _open;                               //�ð�                 
	char opensign                         [  1];	char _opensign;                           //�ð�����ȣ         
	char openchange                       [  6];	char _openchange;                         //�ð��������       
	char low                              [  7];	char _low;                                //��������             
	char dnlmtprice                       [  7];	char _dnlmtprice;                         //���Ѱ�               
	char hotime                           [  8];	char _hotime;                             //ȣ���ð�             
	char offerho                          [  7];	char _offerho;                            //�ŵ��ֿ켱ȣ��       
	char P_offer                          [  7];	char _P_offer;                            //�ŵ�����ȣ��         
	char S_offer                          [  7];	char _S_offer;                            //�ŵ�������ȣ��       
	char S4_offer                         [  7];	char _S4_offer;                           //�ŵ�4����ȣ��        
	char S5_offer                         [  7];	char _S5_offer;                           //�ŵ�5����ȣ��        
	char S6_offer                         [  7];	char _S6_offer;                           //�ŵ�6����ȣ��        
	char S7_offer                         [  7];	char _S7_offer;                           //�ŵ�7����ȣ��        
	char S8_offer                         [  7];	char _S8_offer;                           //�ŵ�8����ȣ��        
	char S9_offer                         [  7];	char _S9_offer;                           //�ŵ�9����ȣ��        
	char S10_offer                        [  7];	char _S10_offer;                          //�ŵ�10����ȣ��       
	char bidho                            [  7];	char _bidho;                              //�ż��ֿ켱ȣ��       
	char P_bid                            [  7];	char _P_bid;                              //�ż�����ȣ��         
	char S_bid                            [  7];	char _S_bid;                              //�ż�������ȣ��       
	char S4_bid                           [  7];	char _S4_bid;                             //�ż�4����ȣ��        
	char S5_bid                           [  7];	char _S5_bid;                             //�ż�5����ȣ��        
	char S6_bid                           [  7];	char _S6_bid;                             //�ż�6����ȣ��        
	char S7_bid                           [  7];	char _S7_bid;                             //�ż�7����ȣ��        
	char S8_bid                           [  7];	char _S8_bid;                             //�ż�8����ȣ��        
	char S9_bid                           [  7];	char _S9_bid;                             //�ż�9����ȣ��        
	char S10_bid                          [  7];	char _S10_bid;                            //�ż�10����ȣ��       
	char offerrem                         [  9];	char _offerrem;                           //�ŵ��ֿ켱�ܷ�       
	char P_offerrem                       [  9];	char _P_offerrem;                         //�ŵ������ܷ�         
	char S_offerrem                       [  9];	char _S_offerrem;                         //�ŵ��������ܷ�       
	char S4_offerrem                      [  9];	char _S4_offerrem;                        //�ŵ�4�����ܷ�        
	char S5_offerrem                      [  9];	char _S5_offerrem;                        //�ŵ�5�����ܷ�        
	char S6_offerrem                      [  9];	char _S6_offerrem;                        //�ŵ�6�����ܷ�        
	char S7_offerrem                      [  9];	char _S7_offerrem;                        //�ŵ�7�����ܷ�        
	char S8_offerrem                      [  9];	char _S8_offerrem;                        //�ŵ�8�����ܷ�        
	char S9_offerrem                      [  9];	char _S9_offerrem;                        //�ŵ�9�����ܷ�        
	char S10_offerrem                     [  9];	char _S10_offerrem;                       //�ŵ�10�����ܷ�       
	char bidrem                           [  9];	char _bidrem;                             //�ż��ֿ켱�ܷ�       
	char P_bidrem                         [  9];	char _P_bidrem;                           //�ż������ܷ�         
	char S_bidrem                         [  9];	char _S_bidrem;                           //�ż��������ܷ�       
	char S4_bidrem                        [  9];	char _S4_bidrem;                          //�ż�4�����ܷ�        
	char S5_bidrem                        [  9];	char _S5_bidrem;                          //�ż�5�����ܷ�        
	char S6_bidrem                        [  9];	char _S6_bidrem;                          //�ż�6�����ܷ�        
	char S7_bidrem                        [  9];	char _S7_bidrem;                          //�ż�7�����ܷ�        
	char S8_bidrem                        [  9];	char _S8_bidrem;                          //�ż�8�����ܷ�        
	char S9_bidrem                        [  9];	char _S9_bidrem;                          //�ż�9�����ܷ�        
	char S10_bidrem                       [  9];	char _S10_bidrem;                         //�ż�10�����ܷ�       
	char T_offerrem                       [  9];	char _T_offerrem;                         //�Ѹŵ��ܷ�           
	char T_bidrem                         [  9];	char _T_bidrem;                           //�Ѹż��ܷ�           
	char O_offerrem                       [  9];	char _O_offerrem;                         //�ð��ܸŵ��ܷ�       
	char O_bidrem                         [  9];	char _O_bidrem;                           //�ð��ܸż��ܷ�       
	char pivot2upz7                       [  7];	char _pivot2upz7;                         //�Ǻ�2������          
	char pivot1upz7                       [  7];	char _pivot1upz7;                         //�Ǻ�1������          
	char pivotz7                          [  7];	char _pivotz7;                            //�Ǻ���               
	char pivot1dnz7                       [  7];	char _pivot1dnz7;                         //�Ǻ�1������          
	char pivot2dnz7                       [  7];	char _pivot2dnz7;                         //�Ǻ�2������          
	char sosokz6                          [  6];	char _sosokz6;                            //�ڽ����ڽ��ڱ���     
	char jisunamez18                      [ 18];	char _jisunamez18;                        //������               
	char capsizez6                        [  6];	char _capsizez6;                          //�ں��ݱԸ�           
	char output1z16                       [ 16];	char _output1z16;                         //����               
	char marcket1z16                      [ 16];	char _marcket1z16;                        //������ġ1            
	char marcket2z16                      [ 16];	char _marcket2z16;                        //������ġ2            
	char marcket3z16                      [ 16];	char _marcket3z16;                        //������ġ3            
	char marcket4z16                      [ 16];	char _marcket4z16;                        //������ġ4            
	char marcket5z16                      [ 16];	char _marcket5z16;                        //������ġ5            
	char marcket6z16                      [ 16];	char _marcket6z16;                        //������ġ6            
	char cbtext                           [  6];	char _cbtext;                             //CB����               
	char parvalue                         [  7];	char _parvalue;                           //�׸鰡               
	char prepricetitlez12                 [ 12];	char _prepricetitlez12;                   //��������Ÿ��Ʋ       
	char prepricez7                       [  7];	char _prepricez7;                         //��������             
	char subprice                         [  7];	char _subprice;                           //��밡               
	char gongpricez7                      [  7];	char _gongpricez7;                        //����               
	char high5                            [  7];	char _high5;                              //5�ϰ�              
	char low5                             [  7];	char _low5;                               //5������              
	char high20                           [  7];	char _high20;                             //20�ϰ�             
	char low20                            [  7];	char _low20;                              //20������             
	char yhigh                            [  7];	char _yhigh;                              //52���ְ�           
	char yhighdate                        [  4];	char _yhighdate;                          //52���ְ���         
	char ylow                             [  7];	char _ylow;                               //52��������           
	char ylowdate                         [  4];	char _ylowdate;                           //52����������         
	char movlistingz8                     [  8];	char _movlistingz8;                       //�����ֽļ�           
	char listing                          [ 12];	char _listing;                            //�����ֽļ�           
	char totpricez9                       [  9];	char _totpricez9;                         //�ð��Ѿ�             
	char tratimez5                        [  5];	char _tratimez5;                          //�ð�                 
	char off_tra1                         [  6];	char _off_tra1;                           //�ŵ��ŷ���1          
	char bid_tra1                         [  6];	char _bid_tra1;                           //�ż��ŷ���1          
	char N_offvolume1                     [  9];	char _N_offvolume1;                       //�ŵ��ŷ���1          
	char N_bidvolume1                     [  9];	char _N_bidvolume1;                       //�ż��ŷ���1          
	char off_tra2                         [  6];	char _off_tra2;                           //�ŵ��ŷ���2          
	char bid_tra2                         [  6];	char _bid_tra2;                           //�ż��ŷ���2          
	char N_offvolume2                     [  9];	char _N_offvolume2;                       //�ŵ��ŷ���2          
	char N_bidvolume2                     [  9];	char _N_bidvolume2;                       //�ż��ŷ���2          
	char off_tra3                         [  6];	char _off_tra3;                           //�ŵ��ŷ���3          
	char bid_tra3                         [  6];	char _bid_tra3;                           //�ż��ŷ���3          
	char N_offvolume3                     [  9];	char _N_offvolume3;                       //�ŵ��ŷ���3          
	char N_bidvolume3                     [  9];	char _N_bidvolume3;                       //�ż��ŷ���3          
	char off_tra4                         [  6];	char _off_tra4;                           //�ŵ��ŷ���4          
	char bid_tra4                         [  6];	char _bid_tra4;                           //�ż��ŷ���4          
	char N_offvolume4                     [  9];	char _N_offvolume4;                       //�ŵ��ŷ���4          
	char N_bidvolume4                     [  9];	char _N_bidvolume4;                       //�ż��ŷ���4          
	char off_tra5                         [  6];	char _off_tra5;                           //�ŵ��ŷ���5          
	char bid_tra5                         [  6];	char _bid_tra5;                           //�ż��ŷ���5          
	char N_offvolume5                     [  9];	char _N_offvolume5;                       //�ŵ��ŷ���5          
	char N_bidvolume5                     [  9];	char _N_bidvolume5;                       //�ż��ŷ���5          
	char N_offvolall                      [  9];	char _N_offvolall;                        //�ŵ��ܱ��ΰŷ���     
	char N_bidvolall                      [  9];	char _N_bidvolall;                        //�ż��ܱ��ΰŷ���     
	char fortimez6                        [  6];	char _fortimez6;                          //�ܱ��νð�           
	char forratez5                        [  5];	char _forratez5;                          //�ܱ���������         
	char settdatez4                       [  4];	char _settdatez4;                         //������               
	char cratez5                          [  5];	char _cratez5;                            //�ܰ����(%)          
	char yudatez4                         [  4];	char _yudatez4;                           //���������           
	char mudatez4                         [  4];	char _mudatez4;                           //���������           
	char yuratez5                         [  5];	char _yuratez5;                           //�����������         
	char muratez5                         [  5];	char _muratez5;                           //�����������         
	char formovolz10                      [ 10];	char _formovolz10;                        //�ܱ��κ����ּ�       
	char jasa                             [  1];	char _jasa;                               //�ڻ���               
	char listdatez8                       [  8];	char _listdatez8;                         //������               
	char daeratez5Tc1151OutBlock2;                        [  5];	char _daeratez5;                          //������������
	char daedatez6                        [  6];	char _daedatez6;                          //��������������       
	char clovergb                         [  1];	char _clovergb;                           //����Ŭ�ι�           
	char depositgb                        [  1];	char _depositgb;                          //���ű���             
	char capital                          [  9];	char _capital;                            //�ں���               
	char N_alloffvol                      [  9];	char _N_alloffvol;                        //��ü�ŷ����ŵ���     
	char N_allbidvol                      [  9];	char _N_allbidvol;                        //��ü�ŷ����ż���     
	char hnamez21                         [ 21];	char _hnamez21;                           //�����2              
	char detourgb                         [  1];	char _detourgb;                           //��ȸ���忩��         
	char yuratez6                         [  6];	char _yuratez6;                           //������ȸ����2        
	char sosokz6_1                        [  6];	char _sosokz6_1;                          //�ڽ��Ǳ���           
	char maedatez4                        [  4];	char _maedatez4;                          //������������         
	char lratez5                          [  5];	char _lratez5;                            //������(%)            
	char perz5                            [  5];	char _perz5;                              //PER                  
	char handogb                          [  1];	char _handogb;                            //���񺰽ſ��ѵ�       
	char avgprice                         [  7];	char _avgprice;                           //���߰�               
	char listing2                         [ 12];	char _listing2;                           //�����ֽļ�_��        
	char addlisting                       [ 12];	char _addlisting;                         //�߰������ּ�         
	char gicomment                        [100];	char _gicomment;                          //����comment          
	char prevolume                        [  9];	char _prevolume;                          //���ϰŷ���           
	char presign                          [  1];	char _presign;                            //���ϴ������ȣ     
	char prechange                        [  6];	char _prechange;                          //���ϴ������       
	char yhigh2                           [  7];	char _yhigh2;                             //�����ְ�           
	char yhighdate2                       [  4];	char _yhighdate2;                         //�����ְ���         
	char ylow2                            [  7];	char _ylow2;                              //����������           
	char ylowdate2                        [  4];	char _ylowdate2;                          //������������         
	char forstock                         [ 15];	char _forstock;                           //�ܱ��κ����ֽļ�     
	char forlmtz5                         [  5];	char _forlmtz5;                           //�ܱ����ѵ���(%)      
	char maeunit                          [  5];	char _maeunit;                            //�Ÿż�������         
	char mass_opt                         [  1];	char _mass_opt;                           //����뷮���ⱸ��     
	char largemgb                         [  1];	char _largemgb;                           //�뷮�Ÿű���         
} Tc1101OutBlock;

typedef struct tagc1101OutBlock2    //�����ŷ����ڷ�, [�ݺ�]
{
	char time                             [  8];	char _time;                               //�ð�                 
	char price                            [  7];	char _price;                              //���簡               
	char sign                             [  1];	char _sign;                               //�����ȣ             
	char change                           [  6];	char _change;                             //�����               
	char offer                            [  7];	char _offer;                              //�ŵ�ȣ��             
	char bid                              [  7];	char _bid;                                //�ż�ȣ��             
	char movolume                         [  8];	char _movolume;                           //�����ŷ���           
	char volume                           [  9];	char _volume;                             //�ŷ���               
} Tc1101OutBlock2;

typedef struct tagc1101OutBlock3    //������ǥ
{
	char dongsi                           [  1];	char _dongsi;                             //����ȣ������         
	char jeqprice                         [  7];	char _jeqprice;                           //����ü�ᰡ           
	char jeqsign                          [  1];	char _jeqsign;                            //����ü���ȣ         
	char jeqchange                        [  6];	char _jeqchange;                          //����ü������       
	char jeqchrate                        [  5];	char _jeqchrate;                          //����ü������       
	char jeqvol                           [  9];	char _jeqvol;                             //����ü�����         
	char chkdataz1                        [  1];	char _chkdataz1;                          //ECN������������      
	char ecn_price                        [  9];	char _ecn_price;                          //ECN��������          
	char ecn_sign                         [  1];	char _ecn_sign;                           //ECN��ȣ              
	char ecn_change                       [  9];	char _ecn_change;                         //ECN�����            
	char ecn_chrate                       [  5];	char _ecn_chrate;                         //ECN�����            
	char ecn_volume                       [ 10];	char _ecn_volume;                         //ECNü�����          
	char ecn_jeqsign                      [  1];	char _ecn_jeqsign;                        //ECN��񿹻�ü���ȣ  
	char ecn_jeqchange                    [  6];	char _ecn_jeqchange;                      //ECN��񿹻�ü������ 
	char ecn_jeqchrate                    [  5];	char _ecn_jeqchrate;                      //ECN��񿹻�ü������ 
} Tc1101OutBlock3;

typedef struct tagc1101
{
	Tc1101InBlock                     c1101inblock                          ;  //�⺻�Է� 
	Tc1101OutBlock                    c1101outblock                         ;  //���񸶽�Ÿ�⺻�ڷ� 
	Tc1101OutBlock2                   c1101outblock2                   [ 20];  //�����ŷ����ڷ� , [�ݺ�]
	Tc1101OutBlock3                   c1101outblock3                        ;  //������ǥ 
} Tc1101;

typedef struct tags4101InBlock    //�⺻�Է�
{
	char formlang                         [  1];	char _formlang;                           //�ѿ�����             
	char fuitemz9                         [  9];	char _fuitemz9;                           //�����ڵ�             
} Ts4101InBlock;

typedef struct tags4101OutBlock    //���񸶽�Ÿ�⺻�ڷ�
{
	char fuitem                           [  4];	char _fuitem;                             //�����ڵ�             
	char fuhname                          [ 12];	char _fuhname;                            //�����               
	char fucurr                           [  5];	char _fucurr;                             //���簡               
	char fusign                           [  1];	char _fusign;                             //�����ȣ             
	char fuchange                         [  5];	char _fuchange;                           //�����               
	char fuchrate                         [  5];	char _fuchrate;                           //�����               
	char fubasis                          [  5];	char _fubasis;                            //���̽ý�             
	char futheoryprice                    [  5];	char _futheoryprice;                      //�̷а�               
	char fugrate                          [  5];	char _fugrate;                            //������               
	char fugratio                         [  5];	char _fugratio;                           //������               
	char fuvolall                         [  7];	char _fuvolall;                           //�ŷ���               
	char fuvalall                         [ 12];	char _fuvalall;                           //�����ŷ����(�鸸)   
	char fuopenyak                        [  7];	char _fuopenyak;                          //�̰�����������       
	char fupreopenyak                     [  7];	char _fupreopenyak;                       //�̰�����������       
	char fuhprice                         [  5];	char _fuhprice;                           //���Ѱ�               
	char fuhigh                           [  5];	char _fuhigh;                             //��                 
	char fuopen                           [  5];	char _fuopen;                             //�ð�                 
	char fuopensign                       [  1];	char _fuopensign;                         //�ð�����ȣ         
	char fuopenchange                     [  5];	char _fuopenchange;                       //�ð������         
	char fulow                            [  5];	char _fulow;                              //����                 
	char fulprice                         [  5];	char _fulprice;                           //���Ѱ�               
	char fucbhprice                       [  5];	char _fucbhprice;                         //CB�ߵ�����           
	char fucblprice                       [  5];	char _fucblprice;                         //CB�ߵ�����           
	char fudehprice                       [  5];	char _fudehprice;                         //DEMARK����           
	char fudelprice                       [  5];	char _fudelprice;                         //DEMARK����           
	char fulisthprice                     [  5];	char _fulisthprice;                       //�������ְ�         
	char fulisthdate                      [  8];	char _fulisthdate;                        //�������ְ���         
	char fulistlprice                     [  5];	char _fulistlprice;                       //������������         
	char fulistldate                      [  8];	char _fulistldate;                        //������������         
	char fulastdate                       [  8];	char _fulastdate;                         //�����ŷ���           
	char fujandatecnt                     [  3];	char _fujandatecnt;                       //������               
	char fucdratio                        [  6];	char _fucdratio;                          //������������         
	char fuchetime                        [  8];	char _fuchetime;                          //ȣ���ð�             
	char fuoffer                          [  5];	char _fuoffer;                            //�ŵ��ֿ켱ȣ��       
	char fujoffer                         [  5];	char _fujoffer;                           //�ŵ�����ȣ��         
	char fujjoffer                        [  5];	char _fujjoffer;                          //�ŵ�������ȣ��       
	char fuj4offer                        [  5];	char _fuj4offer;                          //�ŵ�4����ȣ��        
	char fuj5offer                        [  5];	char _fuj5offer;                          //�ŵ�5����ȣ��        
	char fubid                            [  5];	char _fubid;                              //�ż��ֿ켱ȣ��       
	char fujbid                           [  5];	char _fujbid;                             //�ż�����ȣ��         
	char fujjbid                          [  5];	char _fujjbid;                            //�ż�������ȣ��       
	char fuj4bid                          [  5];	char _fuj4bid;                            //�ż�4����ȣ��        
	char fuj5bid                          [  5];	char _fuj5bid;                            //�ż�5����ȣ��        
	char fuofferjan                       [  6];	char _fuofferjan;                         //�ŵ��ֿ켱�ܷ�       
	char fujofferjan                      [  6];	char _fujofferjan;                        //�ŵ������ܷ�         
	char fujjofferjan                     [  6];	char _fujjofferjan;                       //�ŵ��������ܷ�       
	char fuj4offerjan                     [  6];	char _fuj4offerjan;                       //�ŵ�4�����ܷ�        
	char fuj5offerjan                     [  6];	char _fuj5offerjan;                       //�ŵ�5�����ܷ�        
	char fubidjan                         [  6];	char _fubidjan;                           //�ż��ֿ켱�ܷ�       
	char fujbidjan                        [  6];	char _fujbidjan;                          //�ż������ܷ�         
	char fujjbidjan                       [  6];	char _fujjbidjan;                         //�ż��������ܷ�       
	char fuj4bidjan                       [  6];	char _fuj4bidjan;                         //�ż�4�����ܷ�        
	char fuj5bidjan                       [  6];	char _fuj5bidjan;                         //�ż�5�����ܷ�        
	char futofferjan                      [  6];	char _futofferjan;                        //�Ѹŵ��ܷ�           
	char futbidjan                        [  6];	char _futbidjan;                          //�Ѹż��ܷ�           
	char fuoffersu                        [  4];	char _fuoffersu;                          //�ŵ��ֿ켱�Ǽ�       
	char fujoffersu                       [  4];	char _fujoffersu;                         //�ŵ������Ǽ�         
	char fujjoffersu                      [  4];	char _fujjoffersu;                        //�ŵ��������Ǽ�       
	char fuj4offersu                      [  4];	char _fuj4offersu;                        //�ŵ�4�����Ǽ�        
	char fuj5offersu                      [  4];	char _fuj5offersu;                        //�ŵ�5�����Ǽ�        
	char fubidsu                          [  4];	char _fubidsu;                            //�ż��ֿ켱�Ǽ�       
	char fujbidsu                         [  4];	char _fujbidsu;                           //�ż������Ǽ�         
	char fujjbidsu                        [  4];	char _fujjbidsu;                          //�ż��������Ǽ�       
	char fuj4bidsu                        [  4];	char _fuj4bidsu;                          //�ż�4�����Ǽ�        
	char fuj5bidsu                        [  4];	char _fuj5bidsu;                          //�ż�5�����Ǽ�        
	char futoffersu                       [  5];	char _futoffersu;                         //�Ѹŵ��Ǽ�           
	char futbidsu                         [  5];	char _futbidsu;                           //�Ѹż��Ǽ�           
	char fupivot2upz5                     [  5];	char _fupivot2upz5;                       //�Ǻ�2������          
	char fupivot1upz5                     [  5];	char _fupivot1upz5;                       //�Ǻ�1������          
	char fupivotz5                        [  5];	char _fupivotz5;                          //�Ǻ���               
	char fupivot1dnz5                     [  5];	char _fupivot1dnz5;                       //�Ǻ�1������          
	char fupivot2dnz5                     [  5];	char _fupivot2dnz5;                       //�Ǻ�2������          
	char fujgubun                         [  8];	char _fujgubun;                           //CB�ߵ�����           
	char fuspvolall                       [  7];	char _fuspvolall;                         //��������ŷ���       
	char fudivideratio                    [  9];	char _fudivideratio;                      //��������           
	char preclose                         [  5];	char _preclose;                           //��������             
	char fudynhprice                      [  5];	char _fudynhprice;                        //�ǽð����Ѱ�         
	char fudynlprice                      [  5];	char _fudynlprice;                        //�ǽð����Ѱ�         
	char fudynpriceflag                   [  1];	char _fudynpriceflag;                     //�����������ѿ���     
	char fubulkvol                        [  7];	char _fubulkvol;                          //���ǰŷ���           
	char exlmtgb                          [  1];	char _exlmtgb;                            //����Ȯ�뿹������     
	char uplmtgb                          [  1];	char _uplmtgb;                            //��������Ȯ����Ѵܰ� 
	char dnlmtgb                          [  1];	char _dnlmtgb;                            //��������Ȯ�����Ѵܰ� 
} Ts4101OutBlock;

typedef struct tags4101OutBlock1    //�ڽ���200����
{
	char fuitem                           [  4];	char _fuitem;                             //�ڽ���200�ڵ�        
	char fucurr                           [  5];	char _fucurr;                             //�ڽ���200����        
	char fusign                           [  1];	char _fusign;                             //�ڽ���200�����ȣ    
	char fuchange                         [  5];	char _fuchange;                           //�ڽ���200�����      
	char fuchrate                         [  5];	char _fuchrate;                           //�ڽ���200�����      
} Ts4101OutBlock1;

typedef struct tags4101OutBlock2    //�����ŷ����ڷ�, [�ݺ�]
{
	char fuchetime                        [  8];	char _fuchetime;                          //�ð�                 
	char fucurr                           [  5];	char _fucurr;                             //���簡               
	char fusign                           [  1];	char _fusign;                             //�����ȣ             
	char fuchange                         [  5];	char _fuchange;                           //�����               
	char fuoffer                          [  5];	char _fuoffer;                            //�ŵ�ȣ��             
	char fubid                            [  5];	char _fubid;                              //�ż�ȣ��             
	char fuvol                            [  6];	char _fuvol;                              //�ŷ���               
	char fuvolall                         [  7];	char _fuvolall;                           //�����ŷ���           
	char fuopenyak                        [  7];	char _fuopenyak;                          //�̰�������           
} Ts4101OutBlock2;

typedef struct tags4101OutBlock3    //�ð��뺰��������Ȳ�ֱٸ޸�, [�ݺ�]
{
	char titlez6                          [  6];	char _titlez6;                            //TITLE                
	char amesuvalpure                     [  9];	char _amesuvalpure;                       //���ż�               
	char cmesuvalpure                     [  9];	char _cmesuvalpure;                       //�ŵ�                 
	char imesuvalpure                     [  9];	char _imesuvalpure;                       //�ż�                 
} Ts4101OutBlock3;

typedef struct tags4101OutBlock4    //�ð��뺰��������Ȳ�ð���, [�ݺ�]
{
	char timez8                           [  8];	char _timez8;                             //�ð���               
	char amesuvalpure                     [  9];	char _amesuvalpure;                       //�ܱ��μ��ż�         
	char cmesuvalpure                     [  9];	char _cmesuvalpure;                       //���Ǽ��ż�           
	char imesuvalpure                     [  9];	char _imesuvalpure;                       //���μ��ż�           
} Ts4101OutBlock4;

typedef struct tags4101OutBlock5    //KOSPI200�ð��Ѿ׻���10����, [�ݺ�]
{
	char code                             [  6];	char _code;                               //�����ڵ�             
	char hname                            [ 13];	char _hname;                              //�����               
	char parvalue                         [  7];	char _parvalue;                           //�׸鰡               
	char price                            [  7];	char _price;                              //���簡               
	char sign                             [  1];	char _sign;                               //�����ȣ             
	char change                           [  6];	char _change;                             //�����               
	char chrate                           [  5];	char _chrate;                             //�����               
} Ts4101OutBlock5;

typedef struct tags4101OutBlock6    //����ü��
{
	char dongsi                           [  1];	char _dongsi;                             //����ȣ������         
	char jeqprice                         [  5];	char _jeqprice;                           //����ü�ᰡ           
	char jeqsign                          [  1];	char _jeqsign;                            //����ü���ȣ         
	char jeqchange                        [  5];	char _jeqchange;                          //����ü������       
	char jeqchrate                        [  5];	char _jeqchrate;                          //����ü������       
} Ts4101OutBlock6;

typedef struct tags4101
{
	Ts4101InBlock                     s4101inblock                          ;  //�⺻�Է� 
	Ts4101OutBlock                    s4101outblock                         ;  //���񸶽�Ÿ�⺻�ڷ� 
	Ts4101OutBlock1                   s4101outblock1                        ;  //�ڽ���200���� 
	Ts4101OutBlock2                   s4101outblock2                   [ 30];  //�����ŷ����ڷ� , [�ݺ�]
	Ts4101OutBlock3                   s4101outblock3                   [  3];  //�ð��뺰��������Ȳ�ֱٸ޸� , [�ݺ�]
	Ts4101OutBlock4                   s4101outblock4                   [ 20];  //�ð��뺰��������Ȳ�ð��� , [�ݺ�]
	Ts4101OutBlock5                   s4101outblock5                   [ 10];  //KOSPI200�ð��Ѿ׻���10���� , [�ݺ�]
	Ts4101OutBlock6                   s4101outblock6                        ;  //����ü�� 
} Ts4101;


typedef struct tagc4113InBlock    //�Էµ���Ÿ
{
	char fuitemz9                         [  9];	char _fuitemz9;                           //�Է��ڵ�             
} Tc4113InBlock;

typedef struct tagc4113OutKospi200    //c4113OutKospi200
{
	char fuitem                           [  4];	char _fuitem;                             //�����ڵ�             
	char fucurr                           [  5];	char _fucurr;                             //��������             
	char fusign                           [  1];	char _fusign;                             //���Ϻ��ȣ           
	char fuchange                         [  5];	char _fuchange;                           //���Ϻ�               
	char fuopen                           [  5];	char _fuopen;                             //�ð�                 
	char fuhigh                           [  5];	char _fuhigh;                             //��                 
	char fulow                            [  5];	char _fulow;                              //����                 
	char fuvolall                         [  7];	char _fuvolall;                           //�ŷ���               
} Tc4113OutKospi200;

typedef struct tagc4113OutSMaster    //c4113OutSMaster
{
	char fuitemz8                         [  8];	char _fuitemz8;                           //�����ڵ�             
	char fuspcurr                         [  6];	char _fuspcurr;                           //����                 
	char fuspsign                         [  1];	char _fuspsign;                           //���Ϻ��ȣ           
	char fuspchange                       [  5];	char _fuspchange;                         //���Ϻ�               
	char fuspchrate                       [  5];	char _fuspchrate;                         //�����               
	char fuspopen                         [  6];	char _fuspopen;                           //�ð�                 
	char fusphigh                         [  6];	char _fusphigh;                           //��                 
	char fusplow                          [  6];	char _fusplow;                            //����                 
	char fuspvolall                       [  7];	char _fuspvolall;                         //�ŷ���               
	char fuspvalall                       [ 12];	char _fuspvalall;                         //�����ŷ����(�鸸��) 
	char fuspcurr1                        [  5];	char _fuspcurr1;                          //����������(�ٿ���)   
	char fuspcurr2                        [  5];	char _fuspcurr2;                          //����������(������)   
	char fudynhprice                      [  6];	char _fudynhprice;                        //�ǽð����Ѱ�         
	char fudynlprice                      [  6];	char _fudynlprice;                        //�ǽð����Ѱ�         
	char fudynpriceflag                   [  1];	char _fudynpriceflag;                     //�����������ѿ���     
} Tc4113OutSMaster;

typedef struct tagc4113OutBlock1    //�ڽ��Ǽ���Master1
{
	char fuitem                           [  4];	char _fuitem;                             //�����ڵ�             
	char fuchetime                        [  8];	char _fuchetime;                          //ü��ð�             
	char fuhname                          [ 12];	char _fuhname;                            //�ѱ۸�               
	char fucurr                           [  5];	char _fucurr;                             //���簡               
	char fusign                           [  1];	char _fusign;                             //���ϴ���ȣ         
	char fuchange                         [  5];	char _fuchange;                           //���ϴ��             
	char fuchrate                         [  5];	char _fuchrate;                           //�����               
	char fubasis                          [  5];	char _fubasis;                            //���̽ý�             
	char futheoryprice                    [  5];	char _futheoryprice;                      //�̷а�               
	char fugrate                          [  5];	char _fugrate;                            //������               
	char fugratio                         [  5];	char _fugratio;                           //������               
	char fuvolall                         [  7];	char _fuvolall;                           //����ü�����         
	char fuvalall                         [ 12];	char _fuvalall;                           //�����ŷ����(�鸸��) 
	char fuopenyak                        [  7];	char _fuopenyak;                          //�̰�����������       
	char fupreopenyak                     [  7];	char _fupreopenyak;                       //�̰�����������       
	char fujgubun                         [  8];	char _fujgubun;                           //����               
	char fuopen                           [  5];	char _fuopen;                             //�ð�                 
	char fuhigh                           [  5];	char _fuhigh;                             //��                 
	char fulow                            [  5];	char _fulow;                              //����                 
	char fudynhprice                      [  5];	char _fudynhprice;                        //�ǽð����Ѱ�         
	char fudynlprice                      [  5];	char _fudynlprice;                        //�ǽð����Ѱ�         
	char fudynpriceflag                   [  1];	char _fudynpriceflag;                     //�����������ѿ���     
} Tc4113OutBlock1;

typedef struct tagc4113OutBlock2    //�ڽ��Ǽ���Master2
{
	char fuitem                           [  4];	char _fuitem;                             //�����ڵ�             
	char fuchetime                        [  8];	char _fuchetime;                          //ü��ð�             
	char fuhname                          [ 12];	char _fuhname;                            //�ѱ۸�               
	char fucurr                           [  5];	char _fucurr;                             //���簡               
	char fusign                           [  1];	char _fusign;                             //���ϴ���ȣ         
	char fuchange                         [  5];	char _fuchange;                           //���ϴ��             
	char fuchrate                         [  5];	char _fuchrate;                           //�����               
	char fubasis                          [  5];	char _fubasis;                            //���̽ý�             
	char futheoryprice                    [  5];	char _futheoryprice;                      //�̷а�               
	char fugrate                          [  5];	char _fugrate;                            //������               
	char fugratio                         [  5];	char _fugratio;                           //������               
	char fuvolall                         [  7];	char _fuvolall;                           //����ü�����         
	char fuvalall                         [ 12];	char _fuvalall;                           //�����ŷ����(�鸸��) 
	char fuopenyak                        [  7];	char _fuopenyak;                          //�̰�����������       
	char fupreopenyak                     [  7];	char _fupreopenyak;                       //�̰�����������       
	char fujgubun                         [  8];	char _fujgubun;                           //����               
	char fuopen                           [  5];	char _fuopen;                             //�ð�                 
	char fuhigh                           [  5];	char _fuhigh;                             //��                 
	char fulow                            [  5];	char _fulow;                              //����                 
	char fudynhprice                      [  5];	char _fudynhprice;                        //�ǽð����Ѱ�         
	char fudynlprice                      [  5];	char _fudynlprice;                        //�ǽð����Ѱ�         
	char fudynpriceflag                   [  1];	char _fudynpriceflag;                     //�����������ѿ���     
} Tc4113OutBlock2;

typedef struct tagc4113OutHoga1    //�ڽ��Ǽ���ȣ��1
{
	char fuitem                           [  4];	char _fuitem;                             //�����ڵ�             
	char fuhotime                         [  8];	char _fuhotime;                           //ȣ���ð�             
	char fuoffer                          [  5];	char _fuoffer;                            //�ŵ��켱ȣ��         
	char fujoffer                         [  5];	char _fujoffer;                           //�����ŵ�ȣ��         
	char fujjoffer                        [  5];	char _fujjoffer;                          //�������ŵ�ȣ��       
	char fuj4offer                        [  5];	char _fuj4offer;                          //4�����ŵ�ȣ��        
	char fuj5offer                        [  5];	char _fuj5offer;                          //5�����ŵ�ȣ��        
	char fuofferjan                       [  6];	char _fuofferjan;                         //�ŵ��ܷ�             
	char fujofferjan                      [  6];	char _fujofferjan;                        //�����ŵ�ȣ���ܷ�     
	char fujjofferjan                     [  6];	char _fujjofferjan;                       //�������ŵ�ȣ���ܷ�   
	char fuj4offerjan                     [  6];	char _fuj4offerjan;                       //4�����ŵ�ȣ���ܷ�    
	char fuj5offerjan                     [  6];	char _fuj5offerjan;                       //5�����ŵ�ȣ���ܷ�    
	char fubid                            [  5];	char _fubid;                              //�ż��켱ȣ��         
	char fujbid                           [  5];	char _fujbid;                             //�����ż�ȣ��         
	char fujjbid                          [  5];	char _fujjbid;                            //�������ż�ȣ��       
	char fuj4bid                          [  5];	char _fuj4bid;                            //4�����ż�ȣ��        
	char fuj5bid                          [  5];	char _fuj5bid;                            //5�����ż�ȣ��        
	char fubidjan                         [  6];	char _fubidjan;                           //�ż��ܷ�             
	char fujbidjan                        [  6];	char _fujbidjan;                          //�����ż�ȣ���ܷ�     
	char fujjbidjan                       [  6];	char _fujjbidjan;                         //�������ż�ȣ���ܷ�   
	char fuj4bidjan                       [  6];	char _fuj4bidjan;                         //4�����ż�ȣ���ܷ�    
	char fuj5bidjan                       [  6];	char _fuj5bidjan;                         //5�����ż�ȣ���ܷ�    
	char futofferjan                      [  6];	char _futofferjan;                        //�Ѹŵ��ܷ�           
	char futbidjan                        [  6];	char _futbidjan;                          //�Ѹż��ܷ�           
	char fuoffersu                        [  4];	char _fuoffersu;                          //�ŵ��ֿ켱�Ǽ�       
	char fujoffersu                       [  4];	char _fujoffersu;                         //�ŵ������Ǽ�         
	char fujjoffersu                      [  4];	char _fujjoffersu;                        //�ŵ��������Ǽ�       
	char fuj4offersu                      [  4];	char _fuj4offersu;                        //�ŵ�4�����Ǽ�        
	char fuj5offersu                      [  4];	char _fuj5offersu;                        //�ŵ�5�����Ǽ�        
	char fubidsu                          [  4];	char _fubidsu;                            //�ż��ֿ켱�Ǽ�       
	char fujbidsu                         [  4];	char _fujbidsu;                           //�ż������Ǽ�         
	char fujjbidsu                        [  4];	char _fujjbidsu;                          //�ż��������Ǽ�       
	char fuj4bidsu                        [  4];	char _fuj4bidsu;                          //�ż�4�����Ǽ�        
	char fuj5bidsu                        [  4];	char _fuj5bidsu;                          //�ż�5�����Ǽ�        
	char futoffersu                       [  5];	char _futoffersu;                         //�Ѹŵ��Ǽ�           
	char futbidsu                         [  5];	char _futbidsu;                           //�Ѹż��Ǽ�           
	char fuhname                          [ 12];	char _fuhname;                            //�ѱ۸�               
} Tc4113OutHoga1;

typedef struct tagc4113OutHoga2    //�ڽ��Ǽ���ȣ��2
{
	char fuitem                           [  4];	char _fuitem;                             //�����ڵ�             
	char fuhotime                         [  8];	char _fuhotime;                           //ȣ���ð�             
	char fuoffer                          [  5];	char _fuoffer;                            //�ŵ��켱ȣ��         
	char fujoffer                         [  5];	char _fujoffer;                           //�����ŵ�ȣ��         
	char fujjoffer                        [  5];	char _fujjoffer;                          //�������ŵ�ȣ��       
	char fuj4offer                        [  5];	char _fuj4offer;                          //4�����ŵ�ȣ��        
	char fuj5offer                        [  5];	char _fuj5offer;                          //5�����ŵ�ȣ��        
	char fuofferjan                       [  6];	char _fuofferjan;                         //�ŵ��ܷ�             
	char fujofferjan                      [  6];	char _fujofferjan;                        //�����ŵ�ȣ���ܷ�     
	char fujjofferjan                     [  6];	char _fujjofferjan;                       //�������ŵ�ȣ���ܷ�   
	char fuj4offerjan                     [  6];	char _fuj4offerjan;                       //4�����ŵ�ȣ���ܷ�    
	char fuj5offerjan                     [  6];	char _fuj5offerjan;                       //5�����ŵ�ȣ���ܷ�    
	char fubid                            [  5];	char _fubid;                              //�ż��켱ȣ��         
	char fujbid                           [  5];	char _fujbid;                             //�����ż�ȣ��         
	char fujjbid                          [  5];	char _fujjbid;                            //�������ż�ȣ��       
	char fuj4bid                          [  5];	char _fuj4bid;                            //4�����ż�ȣ��        
	char fuj5bid                          [  5];	char _fuj5bid;                            //5�����ż�ȣ��        
	char fubidjan                         [  6];	char _fubidjan;                           //�ż��ܷ�             
	char fujbidjan                        [  6];	char _fujbidjan;                          //�����ż�ȣ���ܷ�     
	char fujjbidjan                       [  6];	char _fujjbidjan;                         //�������ż�ȣ���ܷ�   
	char fuj4bidjan                       [  6];	char _fuj4bidjan;                         //4�����ż�ȣ���ܷ�    
	char fuj5bidjan                       [  6];	char _fuj5bidjan;                         //5�����ż�ȣ���ܷ�    
	char futofferjan                      [  6];	char _futofferjan;                        //�Ѹŵ��ܷ�           
	char futbidjan                        [  6];	char _futbidjan;                          //�Ѹż��ܷ�           
	char fuoffersu                        [  4];	char _fuoffersu;                          //�ŵ��ֿ켱�Ǽ�       
	char fujoffersu                       [  4];	char _fujoffersu;                         //�ŵ������Ǽ�         
	char fujjoffersu                      [  4];	char _fujjoffersu;                        //�ŵ��������Ǽ�       
	char fuj4offersu                      [  4];	char _fuj4offersu;                        //�ŵ�4�����Ǽ�        
	char fuj5offersu                      [  4];	char _fuj5offersu;                        //�ŵ�5�����Ǽ�        
	char fubidsu                          [  4];	char _fubidsu;                            //�ż��ֿ켱�Ǽ�       
	char fujbidsu                         [  4];	char _fujbidsu;                           //�ż������Ǽ�         
	char fujjbidsu                        [  4];	char _fujjbidsu;                          //�ż��������Ǽ�       
	char fuj4bidsu                        [  4];	char _fuj4bidsu;                          //�ż�4�����Ǽ�        
	char fuj5bidsu                        [  4];	char _fuj5bidsu;                          //�ż�5�����Ǽ�        
	char futoffersu                       [  5];	char _futoffersu;                         //�Ѹŵ��Ǽ�           
	char futbidsu                         [  5];	char _futbidsu;                           //�Ѹż��Ǽ�           
	char fuhname                          [ 12];	char _fuhname;                            //�ѱ۸�               
} Tc4113OutHoga2;

typedef struct tagc4113OutHoga3    //�ڽ��Ǽ�����������ȣ��3
{
	char fuspfuitem                       [  8];	char _fuspfuitem;                         //�����ڵ�             
	char fusphname                        [ 14];	char _fusphname;                          //�ѱ۸�               
	char fusphotime                       [  8];	char _fusphotime;                         //ȣ���ð�             
	char fuspoffer                        [  6];	char _fuspoffer;                          //�ŵ��켱ȣ��         
	char fuspjoffer                       [  6];	char _fuspjoffer;                         //�����ŵ�ȣ��         
	char fuspjjoffer                      [  6];	char _fuspjjoffer;                        //�������ŵ�ȣ��       
	char fuspj4offer                      [  6];	char _fuspj4offer;                        //4�����ŵ�ȣ��        
	char fuspj5offer                      [  6];	char _fuspj5offer;                        //5�����ŵ�ȣ��        
	char fuspofferjan                     [  6];	char _fuspofferjan;                       //�ŵ��ܷ�             
	char fuspjofferjan                    [  6];	char _fuspjofferjan;                      //�����ŵ�ȣ���ܷ�     
	char fuspjjofferjan                   [  6];	char _fuspjjofferjan;                     //�������ŵ�ȣ���ܷ�   
	char fuspj4offerjan                   [  6];	char _fuspj4offerjan;                     //4�����ŵ�ȣ���ܷ�    
	char fuspj5offerjan                   [  6];	char _fuspj5offerjan;                     //5�����ŵ�ȣ���ܷ�    
	char fuspbid                          [  6];	char _fuspbid;                            //�ż��켱ȣ��         
	char fuspjbid                         [  6];	char _fuspjbid;                           //�����ż�ȣ��         
	char fuspjjbid                        [  6];	char _fuspjjbid;                          //�������ż�ȣ��       
	char fuspj4bid                        [  6];	char _fuspj4bid;                          //4�����ż�ȣ��        
	char fuspj5bid                        [  6];	char _fuspj5bid;                          //5�����ż�ȣ��        
	char fuspbidjan                       [  6];	char _fuspbidjan;                         //�ż��ܷ�             
	char fuspjbidjan                      [  6];	char _fuspjbidjan;                        //�����ż�ȣ���ܷ�     
	char fuspjjbidjan                     [  6];	char _fuspjjbidjan;                       //�������ż�ȣ���ܷ�   
	char fuspj4bidjan                     [  6];	char _fuspj4bidjan;                       //4�����ż�ȣ���ܷ�    
	char fuspj5bidjan                     [  6];	char _fuspj5bidjan;                       //5�����ż�ȣ���ܷ�    
	char fusptofferjan                    [  6];	char _fusptofferjan;                      //�Ѹŵ��ܷ�           
	char fusptbidjan                      [  6];	char _fusptbidjan;                        //�Ѹż��ܷ�           
	char fuspoffersu                      [  4];	char _fuspoffersu;                        //�ŵ��ֿ켱�Ǽ�       
	char fuspjoffersu                     [  4];	char _fuspjoffersu;                       //�ŵ������Ǽ�         
	char fuspjjoffersu                    [  4];	char _fuspjjoffersu;                      //�ŵ��������Ǽ�       
	char fuspj4offersu                    [  4];	char _fuspj4offersu;                      //�ŵ�4�����Ǽ�        
	char fuspj5offersu                    [  4];	char _fuspj5offersu;                      //�ŵ�5�����Ǽ�        
	char fuspbidsu                        [  4];	char _fuspbidsu;                          //�ż��ֿ켱�Ǽ�       
	char fuspjbidsu                       [  4];	char _fuspjbidsu;                         //�ż������Ǽ�         
	char fuspjjbidsu                      [  4];	char _fuspjjbidsu;                        //�ż��������Ǽ�       
	char fuspj4bidsu                      [  4];	char _fuspj4bidsu;                        //�ż�4�����Ǽ�        
	char fuspj5bidsu                      [  4];	char _fuspj5bidsu;                        //�ż�5�����Ǽ�        
	char fusptoffersu                     [  5];	char _fusptoffersu;                       //�Ѹŵ��Ǽ�           
	char fusptbidsu                       [  5];	char _fusptbidsu;                         //�Ѹż��Ǽ�           
} Tc4113OutHoga3;

typedef struct tagc4113OutFuteq1    //��������ü��1
{
	char dongsi                           [  1];	char _dongsi;                             //����ȣ������         
	char jeqprice                         [  5];	char _jeqprice;                           //����ü�ᰡ           
	char jeqsign                          [  1];	char _jeqsign;                            //����ü���ȣ         
	char jeqchange                        [  5];	char _jeqchange;                          //����ü������       
	char jeqchrate                        [  5];	char _jeqchrate;                          //����ü������       
} Tc4113OutFuteq1;

typedef struct tagc4113OutFuteq2    //��������ü��2
{
	char dongsi                           [  1];	char _dongsi;                             //����ȣ������         
	char jeqprice                         [  5];	char _jeqprice;                           //����ü�ᰡ           
	char jeqsign                          [  1];	char _jeqsign;                            //����ü���ȣ         
	char jeqchange                        [  5];	char _jeqchange;                          //����ü������       
	char jeqchrate                        [  5];	char _jeqchrate;                          //����ü������       
} Tc4113OutFuteq2;

typedef struct tagc4113
{
	Tc4113InBlock                     c4113inblock                          ;  //�Էµ���Ÿ 
	Tc4113OutKospi200                 c4113outkospi200                      ;  //c4113OutKospi200 
	Tc4113OutSMaster                  c4113outsmaster                       ;  //c4113OutSMaster 
	Tc4113OutBlock1                   c4113outblock1                        ;  //�ڽ��Ǽ���Master1 
	Tc4113OutBlock2                   c4113outblock2                        ;  //�ڽ��Ǽ���Master2 
	Tc4113OutHoga1                    c4113outhoga1                         ;  //�ڽ��Ǽ���ȣ��1 
	Tc4113OutHoga2                    c4113outhoga2                         ;  //�ڽ��Ǽ���ȣ��2 
	Tc4113OutHoga3                    c4113outhoga3                         ;  //�ڽ��Ǽ�����������ȣ��3 
	Tc4113OutFuteq1                   c4113outfuteq1                        ;  //��������ü��1 
	Tc4113OutFuteq2                   c4113outfuteq2                        ;  //��������ü��2 
} Tc4113;

typedef struct tags4201InBlock    //�⺻�Է�
{
	char formlang                         [  1];	char _formlang;                           //�ѿ�����             
	char opitemz9                         [  9];	char _opitemz9;                           //�����ڵ�             
} Ts4201InBlock;

typedef struct tags4201OutBlock    //���񸶽�Ÿ�⺻�ڷ�
{
	char opitem                           [  8];	char _opitem;                             //�����ڵ�             
	char ophname                          [ 14];	char _ophname;                            //�����               
	char opcurr                           [  5];	char _opcurr;                             //���簡               
	char opsign                           [  1];	char _opsign;                             //�����ȣ             
	char opchange                         [  5];	char _opchange;                           //�����               
	char opchrate                         [  5];	char _opchrate;                           //�����               
	char opopen                           [  5];	char _opopen;                             //�ð�                 
	char ophigh                           [  5];	char _ophigh;                             //��                 
	char oplow                            [  5];	char _oplow;                              //����                 
	char optheoryprice                    [  5];	char _optheoryprice;                      //�̷а�               
	char opvolallz8                       [  8];	char _opvolallz8;                         //�ŷ���               
	char opvalall                         [ 12];	char _opvalall;                           //�����ŷ����(�鸸)   
	char opopenyak                        [  7];	char _opopenyak;                          //�̰�����������       
	char oppreopenyak                     [  7];	char _oppreopenyak;                       //�̰�����������       
	char oplisthdatez11                   [ 11];	char _oplisthdatez11;                     //�������ְ���         
	char oplistldatez11                   [ 11];	char _oplistldatez11;                     //������������         
	char oplistdate                       [  8];	char _oplistdate;                         //�ŷ�������           
	char oplastdate                       [  8];	char _oplastdate;                         //�����ŷ���           
	char opjandatecnt                     [  4];	char _opjandatecnt;                       //������               
	char ophprice                         [  5];	char _ophprice;                           //���Ѱ�               
	char oplprice                         [  5];	char _oplprice;                           //���Ѱ�               
	char opgrate                          [  5];	char _opgrate;                            //������               
	char opimpv                           [  5];	char _opimpv;                             //���纯����           
	char oppastv90                        [  5];	char _oppastv90;                          //���ź�����90         
	char opdelta                          [  8];	char _opdelta;                            //��Ÿ����             
	char opgmma                           [  8];	char _opgmma;                             //��������             
	char opvega                           [  8];	char _opvega;                             //����������           
	char optheta                          [  8];	char _optheta;                            //��Ÿ�ð�             
	char oprho                            [  8];	char _oprho;                              //��������             
	char opcdratio                        [  6];	char _opcdratio;                          //������               
	char opdivideratio                    [  9];	char _opdivideratio;                      //��������           
	char opchetime                        [  8];	char _opchetime;                          //ȣ���ð�             
	char opoffer                          [  5];	char _opoffer;                            //�ŵ��ֿ켱ȣ��       
	char opjoffer                         [  5];	char _opjoffer;                           //�ŵ�����ȣ��         
	char opjjoffer                        [  5];	char _opjjoffer;                          //�ŵ�������ȣ��       
	char opj4offer                        [  5];	char _opj4offer;                          //�ŵ�4����ȣ��        
	char opj5offer                        [  5];	char _opj5offer;                          //�ŵ�5����ȣ��        
	char opbid                            [  5];	char _opbid;                              //�ż��ֿ켱ȣ��       
	char opjbid                           [  5];	char _opjbid;                             //�ż�����ȣ��         
	char opjjbid                          [  5];	char _opjjbid;                            //�ż�������ȣ��       
	char opj4bid                          [  5];	char _opj4bid;                            //�ż�4����ȣ��        
	char opj5bid                          [  5];	char _opj5bid;                            //�ż�5����ȣ��        
	char opofferjan                       [  7];	char _opofferjan;                         //�ŵ��ֿ켱�ܷ�       
	char opjofferjan                      [  7];	char _opjofferjan;                        //�ŵ������ܷ�         
	char opjjofferjan                     [  7];	char _opjjofferjan;                       //�ŵ��������ܷ�       
	char opj4offerjan                     [  7];	char _opj4offerjan;                       //�ŵ�4�����ܷ�        
	char opj5offerjan                     [  7];	char _opj5offerjan;                       //�ŵ�5�����ܷ�        
	char opbidjan                         [  7];	char _opbidjan;                           //�ż��ֿ켱�ܷ�       
	char opjbidjan                        [  7];	char _opjbidjan;                          //�ż������ܷ�         
	char opjjbidjan                       [  7];	char _opjjbidjan;                         //�ż��������ܷ�       
	char opj4bidjan                       [  7];	char _opj4bidjan;                         //�ż�4�����ܷ�        
	char opj5bidjan                       [  7];	char _opj5bidjan;                         //�ż�5�����ܷ�        
	char optofferjan                      [  7];	char _optofferjan;                        //�Ѹŵ��ܷ�           
	char optbidjan                        [  7];	char _optbidjan;                          //�Ѹż��ܷ�           
	char opoffersu                        [  4];	char _opoffersu;                          //�ŵ��ֿ켱�Ǽ�       
	char opjoffersu                       [  4];	char _opjoffersu;                         //�ŵ������Ǽ�         
	char opjjoffersu                      [  4];	char _opjjoffersu;                        //�ŵ��������Ǽ�       
	char opj4offersu                      [  4];	char _opj4offersu;                        //�ŵ�4�����Ǽ�        
	char opj5offersu                      [  4];	char _opj5offersu;                        //�ŵ�5�����Ǽ�        
	char opbidsu                          [  4];	char _opbidsu;                            //�ż��ֿ켱�Ǽ�       
	char opjbidsu                         [  4];	char _opjbidsu;                           //�ż������Ǽ�         
	char opjjbidsu                        [  4];	char _opjjbidsu;                          //�ż��������Ǽ�       
	char opj4bidsu                        [  4];	char _opj4bidsu;                          //�ż�4�����Ǽ�        
	char opj5bidsu                        [  4];	char _opj5bidsu;                          //�ż�5�����Ǽ�        
	char optoffersu                       [  5];	char _optoffersu;                         //�Ѹŵ��Ǽ�           
	char optbidsu                         [  5];	char _optbidsu;                           //�Ѹż��Ǽ�           
	char opjgubun                         [  8];	char _opjgubun;                           //CB�ߵ�����           
	char opopensign                       [  1];	char _opopensign;                         //�ð�����ȣ         
	char opopenchange                     [  5];	char _opopenchange;                       //�ð������         
	char opgratio                         [  5];	char _opgratio;                           //������               
	char preclose                         [  5];	char _preclose;                           //��������             
	char fudynhprice                      [  5];	char _fudynhprice;                        //�ǽð����Ѱ�         
	char fudynlprice                      [  5];	char _fudynlprice;                        //�ǽð����Ѱ�         
	char fudynpriceflag                   [  1];	char _fudynpriceflag;                     //�����������ѿ���     
	char opbulkvol                        [  8];	char _opbulkvol;                          //���ǰŷ���           
	char uplmtgb                          [  1];	char _uplmtgb;                            //��������Ȯ����Ѵܰ� 
	char dnlmtgb                          [  1];	char _dnlmtgb;                            //��������Ȯ�����Ѵܰ� 
} Ts4201OutBlock;

typedef struct tags4201OutBlock1    //�ڽ���200����
{
	char fuitem                           [  4];	char _fuitem;                             //�ڽ���200�ڵ�        
	char fucurr                           [  5];	char _fucurr;                             //�ڽ���200����        
	char fusign                           [  1];	char _fusign;                             //�ڽ���200�����ȣ    
	char fuchange                         [  5];	char _fuchange;                           //�ڽ���200�����      
	char fuchrate                         [  5];	char _fuchrate;                           //�ڽ���200�����      
} Ts4201OutBlock1;

typedef struct tags4201OutBlock2    //�ɼǺ����ŷ����ڷ�, [�ݺ�]
{
	char opchetime                        [  8];	char _opchetime;                          //�ð�                 
	char opcurr                           [  5];	char _opcurr;                             //���簡               
	char opsign                           [  1];	char _opsign;                             //�����ȣ             
	char opchange                         [  5];	char _opchange;                           //�����               
	char opoffer                          [  5];	char _opoffer;                            //�ŵ�ȣ��             
	char opbid                            [  5];	char _opbid;                              //�ż�ȣ��             
	char opvol                            [  6];	char _opvol;                              //�ŷ���               
	char opvolallz8                       [  8];	char _opvolallz8;                         //�����ŷ���           
	char opopenyak                        [  7];	char _opopenyak;                          //�̰�������           
} Ts4201OutBlock2;

typedef struct tags4201OutBlock3    //�����ֱٿ���
{
	char fuitem                           [  4];	char _fuitem;                             //�����ֱٿ����ڵ�     
	char fuitemz9                         [  9];	char _fuitemz9;                           //�����ֱٿ���Ȯ���ڵ� 
	char fuhname                          [ 12];	char _fuhname;                            //�����ֱٿ�����       
	char fucurr                           [  5];	char _fucurr;                             //�����ֱٿ�������     
	char fusign                           [  1];	char _fusign;                             //�����ֱٿ��������ȣ 
	char fuchange                         [  5];	char _fuchange;                           //�����ֱٿ��������   
	char fuchrate                         [  5];	char _fuchrate;                           //�����ֱٿ��������   
	char fuvolall                         [  7];	char _fuvolall;                           //�����ֱٿ����ŷ���   
	char fuvalall                         [ 12];	char _fuvalall;                           //�����ֱٿ��������ŷ����(�鸸) 
	char fuchetime                        [  8];	char _fuchetime;                          //�����ֱٿ���ȣ���ð� 
	char fuoffer                          [  5];	char _fuoffer;                            //�����ֱٿ����ŵ��ֿ켱ȣ�� 
	char fujoffer                         [  5];	char _fujoffer;                           //�����ֱٿ����ŵ�����ȣ�� 
	char fujjoffer                        [  5];	char _fujjoffer;                          //�����ֱٿ����ŵ�������ȣ�� 
	char fuj4offer                        [  5];	char _fuj4offer;                          //�����ֱٿ����ŵ�4����ȣ�� 
	char fuj5offer                        [  5];	char _fuj5offer;                          //�����ֱٿ����ŵ�5����ȣ�� 
	char fubid                            [  5];	char _fubid;                              //�����ֱٿ����ż��ֿ켱ȣ�� 
	char fujbid                           [  5];	char _fujbid;                             //�����ֱٿ����ż�����ȣ�� 
	char fujjbid                          [  5];	char _fujjbid;                            //�����ֱٿ����ż�������ȣ�� 
	char fuj4bid                          [  5];	char _fuj4bid;                            //�����ֱٿ����ż�4����ȣ�� 
	char fuj5bid                          [  5];	char _fuj5bid;                            //�����ֱٿ����ż�5����ȣ�� 
	char fuofferjan                       [  6];	char _fuofferjan;                         //�����ֱٿ����ŵ��ֿ켱�ܷ� 
	char fujofferjan                      [  6];	char _fujofferjan;                        //�����ֱٿ����ŵ������ܷ� 
	char fujjofferjan                     [  6];	char _fujjofferjan;                       //�����ֱٿ����ŵ��������ܷ� 
	char fuj4offerjan                     [  6];	char _fuj4offerjan;                       //�����ֱٿ����ŵ�4�����ܷ� 
	char fuj5offerjan                     [  6];	char _fuj5offerjan;                       //�����ֱٿ����ŵ�5�����ܷ� 
	char fubidjan                         [  6];	char _fubidjan;                           //�����ֱٿ����ż��ֿ켱�ܷ� 
	char fujbidjan                        [  6];	char _fujbidjan;                          //�����ֱٿ����ż������ܷ� 
	char fujjbidjan                       [  6];	char _fujjbidjan;                         //�����ֱٿ����ż��������ܷ� 
	char fuj4bidjan                       [  6];	char _fuj4bidjan;                         //�����ֱٿ����ż�4�����ܷ� 
	char fuj5bidjan                       [  6];	char _fuj5bidjan;                         //�����ֱٿ����ż�5�����ܷ� 
	char futofferjan                      [  6];	char _futofferjan;                        //�����ֱٿ����Ѹŵ��ܷ� 
	char futbidjan                        [  6];	char _futbidjan;                          //�����ֱٿ����Ѹż��ܷ� 
	char fuoffersu                        [  4];	char _fuoffersu;                          //�����ֱٿ����ŵ��ֿ켱�Ǽ� 
	char fujoffersu                       [  4];	char _fujoffersu;                         //�����ֱٿ����ŵ������Ǽ� 
	char fujjoffersu                      [  4];	char _fujjoffersu;                        //�����ֱٿ����ŵ��������Ǽ� 
	char fuj4offersu                      [  4];	char _fuj4offersu;                        //�����ֱٿ����ŵ�4�����Ǽ� 
	char fuj5offersu                      [  4];	char _fuj5offersu;                        //�����ֱٿ����ŵ�5�����Ǽ� 
	char fubidsu                          [  4];	char _fubidsu;                            //�����ֱٿ����ż��ֿ켱�Ǽ� 
	char fujbidsu                         [  4];	char _fujbidsu;                           //�����ֱٿ����ż������Ǽ� 
	char fujjbidsu                        [  4];	char _fujjbidsu;                          //�����ֱٿ����ż��������Ǽ� 
	char fuj4bidsu                        [  4];	char _fuj4bidsu;                          //�����ֱٿ����ż�4�����Ǽ� 
	char fuj5bidsu                        [  4];	char _fuj5bidsu;                          //�����ֱٿ����ż�5�����Ǽ� 
	char futoffersu                       [  5];	char _futoffersu;                         //�����ֱٿ����Ѹŵ��Ǽ� 
	char futbidsu                         [  5];	char _futbidsu;                           //�����ֱٿ����Ѹż��Ǽ� 
	char fuhprice                         [  5];	char _fuhprice;                           //���Ѱ�               
	char fuhigh                           [  5];	char _fuhigh;                             //��                 
	char fuopen                           [  5];	char _fuopen;                             //�ð�                 
	char fuopensign                       [  1];	char _fuopensign;                         //�ð�����ȣ         
	char fuopenchange                     [  5];	char _fuopenchange;                       //�ð������         
	char fulow                            [  5];	char _fulow;                              //����                 
	char fulprice                         [  5];	char _fulprice;                           //���Ѱ�               
	char fupivot2upz5                     [  5];	char _fupivot2upz5;                       //�Ǻ�2������          
	char fupivot1upz5                     [  5];	char _fupivot1upz5;                       //�Ǻ�1������          
	char fupivotz5                        [  5];	char _fupivotz5;                          //�Ǻ���               
	char fupivot1dnz5                     [  5];	char _fupivot1dnz5;                       //�Ǻ�1������          
	char fupivot2dnz5                     [  5];	char _fupivot2dnz5;                       //�Ǻ�2������          
	char fudynhprice                      [  5];	char _fudynhprice;                        //�ǽð����Ѱ�         
	char fudynlprice                      [  5];	char _fudynlprice;                        //�ǽð����Ѱ�         
	char fudynpriceflag                   [  1];	char _fudynpriceflag;                     //�����������ѿ���     
} Ts4201OutBlock3;

typedef struct tags4201OutBlock4    //���������ŷ����ڷ�, [�ݺ�]
{
	char fuchetime                        [  8];	char _fuchetime;                          //�ð�                 
	char fucurr                           [  5];	char _fucurr;                             //���簡               
	char fusign                           [  1];	char _fusign;                             //�����ȣ             
	char fuchange                         [  5];	char _fuchange;                           //�����               
	char fuoffer                          [  5];	char _fuoffer;                            //�ŵ�ȣ��             
	char fubid                            [  5];	char _fubid;                              //�ż�ȣ��             
	char fuvol                            [  6];	char _fuvol;                              //�ŷ���               
	char fuvolall                         [  7];	char _fuvolall;                           //�����ŷ���           
	char fuopenyak                        [  7];	char _fuopenyak;                          //�̰�������           
} Ts4201OutBlock4;

typedef struct tags4201OutBlock5    //�ɼǿ���ü��
{
	char dongsi                           [  1];	char _dongsi;                             //����ȣ������         
	char jeqprice                         [  5];	char _jeqprice;                           //����ü�ᰡ           
	char jeqsign                          [  1];	char _jeqsign;                            //����ü���ȣ         
	char jeqchange                        [  5];	char _jeqchange;                          //����ü������       
	char jeqchrate                        [  5];	char _jeqchrate;                          //����ü������       
} Ts4201OutBlock5;

typedef struct tags4201OutBlock6    //��������ü��
{
	char dongsi                           [  1];	char _dongsi;                             //����ȣ������         
	char jeqprice                         [  5];	char _jeqprice;                           //����ü�ᰡ           
	char jeqsign                          [  1];	char _jeqsign;                            //����ü���ȣ         
	char jeqchange                        [  5];	char _jeqchange;                          //����ü������       
	char jeqchrate                        [  5];	char _jeqchrate;                          //����ü������       
} Ts4201OutBlock6;

typedef struct tags4201
{
	Ts4201InBlock                     s4201inblock                          ;  //�⺻�Է� 
	Ts4201OutBlock                    s4201outblock                         ;  //���񸶽�Ÿ�⺻�ڷ� 
	Ts4201OutBlock1                   s4201outblock1                        ;  //�ڽ���200���� 
	Ts4201OutBlock2                   s4201outblock2                   [ 20];  //�ɼǺ����ŷ����ڷ� , [�ݺ�]
	Ts4201OutBlock3                   s4201outblock3                        ;  //�����ֱٿ��� 
	Ts4201OutBlock4                   s4201outblock4                   [ 20];  //���������ŷ����ڷ� , [�ݺ�]
	Ts4201OutBlock5                   s4201outblock5                        ;  //�ɼǿ���ü�� 
	Ts4201OutBlock6                   s4201outblock6                        ;  //��������ü�� 
} Ts4201;


typedef struct tagc4801InBlock    //�⺻�Է�
{
	char formlang                         [  1];	char _formlang;                           //�ѿ�����             
	char fuitemz9                         [  9];	char _fuitemz9;                           //�����ڵ�             
} Tc4801InBlock;

typedef struct tagc4801OutBlock    //�ֽļ���MASTER�⺻�ڷ�
{
	char expcode                          [  8];	char _expcode;                            //�����ڵ�             
	char hname                            [ 50];	char _hname;                              //�ѱ۸�               
	char ename                            [ 50];	char _ename;                              //������               
	char sname                            [ 25];	char _sname;                              //�����               
	char baseprice                        [  7];	char _baseprice;                          //���ذ���             
	char hprice                           [  7];	char _hprice;                             //���Ѱ�               
	char lprice                           [  7];	char _lprice;                             //���Ѱ�               
	char preclose                         [  7];	char _preclose;                           //��������             
	char unit                             [ 16];	char _unit;                               //�ŷ�����             
	char openyak                          [  7];	char _openyak;                            //�̰�����������       
	char fusign                           [  1];	char _fusign;                             //���ϴ���ȣ         
	char fuchange                         [  7];	char _fuchange;                           //���ϴ��             
	char fucurr                           [  7];	char _fucurr;                             //���簡               
	char fuopen                           [  7];	char _fuopen;                             //�ð�                 
	char fuhigh                           [  7];	char _fuhigh;                             //��                 
	char fulow                            [  7];	char _fulow;                              //����                 
	char fuvolall                         [  7];	char _fuvolall;                           //����ü�����(���)   
	char fuspvolall                       [  7];	char _fuspvolall;                         //��������ü�����     
	char fuvalall                         [ 12];	char _fuvalall;                           //�����ŷ����(õ��)   
	char hotime                           [  8];	char _hotime;                             //ȣ���ð�             
	char offer                            [  7];	char _offer;                              //�ŵ��켱ȣ��         
	char bid                              [  7];	char _bid;                                //�ż��켱ȣ��         
	char offerjan                         [  6];	char _offerjan;                           //�ŵ��ܷ�             
	char bidjan                           [  6];	char _bidjan;                             //�ż��ܷ�             
	char S2offer                          [  7];	char _S2offer;                            //�ŵ�2��ȣ��          
	char S2bid                            [  7];	char _S2bid;                              //�ż�2��ȣ��          
	char S2offerjan                       [  6];	char _S2offerjan;                         //�ŵ�2���ܷ�          
	char S2bidjan                         [  6];	char _S2bidjan;                           //�ż�2���ܷ�          
	char S3offer                          [  7];	char _S3offer;                            //�ŵ�3��ȣ��          
	char S3bid                            [  7];	char _S3bid;                              //�ż�3��ȣ��          
	char S3offerjan                       [  6];	char _S3offerjan;                         //�ŵ�3���ܷ�          
	char S3bidjan                         [  6];	char _S3bidjan;                           //�ż�3���ܷ�          
	char S4offer                          [  7];	char _S4offer;                            //�ŵ�4��ȣ��          
	char S4bid                            [  7];	char _S4bid;                              //�ż�4��ȣ��          
	char S4offerjan                       [  6];	char _S4offerjan;                         //�ŵ�4���ܷ�          
	char S4bidjan                         [  6];	char _S4bidjan;                           //�ż�4���ܷ�          
	char S5offer                          [  7];	char _S5offer;                            //�ŵ�5��ȣ��          
	char S5bid                            [  7];	char _S5bid;                              //�ż�5��ȣ��          
	char S5offerjan                       [  6];	char _S5offerjan;                         //�ŵ�5���ܷ�          
	char S5bidjan                         [  6];	char _S5bidjan;                           //�ż�5���ܷ�          
	char S6offer                          [  7];	char _S6offer;                            //�ŵ�6��ȣ��          
	char S6bid                            [  7];	char _S6bid;                              //�ż�6��ȣ��          
	char S6offerjan                       [  6];	char _S6offerjan;                         //�ŵ�6���ܷ�          
	char S6bidjan                         [  6];	char _S6bidjan;                           //�ż�6���ܷ�          
	char S7offer                          [  7];	char _S7offer;                            //�ŵ�7��ȣ��          
	char S7bid                            [  7];	char _S7bid;                              //�ż�7��ȣ��          
	char S7offerjan                       [  6];	char _S7offerjan;                         //�ŵ�7���ܷ�          
	char S7bidjan                         [  6];	char _S7bidjan;                           //�ż�7���ܷ�          
	char S8offer                          [  7];	char _S8offer;                            //�ŵ�8��ȣ��          
	char S8bid                            [  7];	char _S8bid;                              //�ż�8��ȣ��          
	char S8offerjan                       [  6];	char _S8offerjan;                         //�ŵ�8���ܷ�          
	char S8bidjan                         [  6];	char _S8bidjan;                           //�ż�8���ܷ�          
	char S9offer                          [  7];	char _S9offer;                            //�ŵ�9��ȣ��          
	char S9bid                            [  7];	char _S9bid;                              //�ż�9��ȣ��          
	char S9offerjan                       [  6];	char _S9offerjan;                         //�ŵ�9���ܷ�          
	char S9bidjan                         [  6];	char _S9bidjan;                           //�ż�9���ܷ�          
	char S0offer                          [  7];	char _S0offer;                            //�ŵ�10��ȣ��         
	char S0bid                            [  7];	char _S0bid;                              //�ż�10��ȣ��         
	char S0offerjan                       [  6];	char _S0offerjan;                         //�ŵ�10���ܷ�         
	char S0bidjan                         [  6];	char _S0bidjan;                           //�ż�10���ܷ�         
	char offersu                          [  4];	char _offersu;                            //�ŵ��Ǽ�             
	char bidsu                            [  4];	char _bidsu;                              //�ż��Ǽ�             
	char S2offersu                        [  4];	char _S2offersu;                          //�ŵ�2���Ǽ�          
	char S2bidsu                          [  4];	char _S2bidsu;                            //�ż�2���Ǽ�          
	char S3offersu                        [  4];	char _S3offersu;                          //�ŵ�3���Ǽ�          
	char S3bidsu                          [  4];	char _S3bidsu;                            //�ż�3���Ǽ�          
	char S4offersu                        [  4];	char _S4offersu;                          //�ŵ�4���Ǽ�          
	char S4bidsu                          [  4];	char _S4bidsu;                            //�ż�4���Ǽ�          
	char S5offersu                        [  4];	char _S5offersu;                          //�ŵ�5���Ǽ�          
	char S5bidsu                          [  4];	char _S5bidsu;                            //�ż�5���Ǽ�          
	char S6offersu                        [  4];	char _S6offersu;                          //�ŵ�6���Ǽ�          
	char S6bidsu                          [  4];	char _S6bidsu;                            //�ż�6���Ǽ�          
	char S7offersu                        [  4];	char _S7offersu;                          //�ŵ�7���Ǽ�          
	char S7bidsu                          [  4];	char _S7bidsu;                            //�ż�7���Ǽ�          
	char S8offersu                        [  4];	char _S8offersu;                          //�ŵ�8���Ǽ�          
	char S8bidsu                          [  4];	char _S8bidsu;                            //�ż�8���Ǽ�          
	char S9offersu                        [  4];	char _S9offersu;                          //�ŵ�9���Ǽ�          
	char S9bidsu                          [  4];	char _S9bidsu;                            //�ż�9���Ǽ�          
	char S0offersu                        [  4];	char _S0offersu;                          //�ŵ�10���Ǽ�         
	char S0bidsu                          [  4];	char _S0bidsu;                            //�ż�10���Ǽ�         
	char tofferjan                        [  6];	char _tofferjan;                          //�Ѹŵ��ܷ�           
	char tobidjan                         [  6];	char _tobidjan;                           //�Ѹż��ܷ�           
	char toffersu                         [  5];	char _toffersu;                           //�Ѹŵ��Ǽ�           
	char tbidsu                           [  5];	char _tbidsu;                             //�Ѹż��Ǽ�           
	char theorytime                       [  6];	char _theorytime;                         //�̷а��ð�           
	char theoryprice                      [  7];	char _theoryprice;                        //�̷а�               
	char fuchrate                         [  5];	char _fuchrate;                           //�����               
	char fupivot2upz7                     [  7];	char _fupivot2upz7;                       //�Ǻ�2������          
	char fupivot1upz7                     [  7];	char _fupivot1upz7;                       //�Ǻ�1������          
	char fupivotz7                        [  7];	char _fupivotz7;                          //�Ǻ���               
	char fupivot1dnz7                     [  7];	char _fupivot1dnz7;                       //�Ǻ�1������          
	char fupivot2dnz7                     [  7];	char _fupivot2dnz7;                       //�Ǻ�2������          
	char fubasis                          [  7];	char _fubasis;                            //���̽ý�             
	char fugrate                          [  7];	char _fugrate;                            //������               
	char fugratio                         [  6];	char _fugratio;                           //������               
	char fupreopenyak                     [  7];	char _fupreopenyak;                       //�̰�����������       
	char fulisthprice                     [  7];	char _fulisthprice;                       //�������ְ�         
	char fulisthdate                      [  8];	char _fulisthdate;                        //�������ְ���         
	char fulistlprice                     [  7];	char _fulistlprice;                       //������������         
	char fulistldate                      [  8];	char _fulistldate;                        //������������         
	char fulastdate                       [  8];	char _fulastdate;                         //�����ŷ���           
	char fujandatecnt                     [  3];	char _fujandatecnt;                       //������               
	char fucdratio                        [  6];	char _fucdratio;                          //������������         
	char fuopenchange                     [  7];	char _fuopenchange;                       //�ð������         
	char fudynhprice                      [  7];	char _fudynhprice;                        //�ǽð����Ѱ�         
	char fudynlprice                      [  7];	char _fudynlprice;                        //�ǽð����Ѱ�         
	char fudynpriceflag                   [  1];	char _fudynpriceflag;                     //�����������ѿ���     
	char exlmtgb                          [  1];	char _exlmtgb;                            //����Ȯ�뿹������     
	char uplmtgb                          [  1];	char _uplmtgb;                            //��������Ȯ����Ѵܰ� 
	char dnlmtgb                          [  1];	char _dnlmtgb;                            //��������Ȯ�����Ѵܰ� 
} Tc4801OutBlock;

typedef struct tagc4801OutBlock1    //�����ڻ�
{
	char shcode                           [  6];	char _shcode;                             //�����ڵ�             
	char hname                            [ 20];	char _hname;                              //�����               
	char price                            [  7];	char _price;                              //���簡               
	char sign                             [  1];	char _sign;                               //�����ȣ             
	char change                           [  6];	char _change;                             //�����               
	char market                           [ 16];	char _market;                             //������               
	char chrate                           [  5];	char _chrate;                             //�����               
	char offer                            [  7];	char _offer;                              //�ŵ�ȣ��             
	char bid                              [  7];	char _bid;                                //�ż�ȣ��             
	char volume                           [  9];	char _volume;                             //�ŷ���               
	char volrate                          [  6];	char _volrate;                            //�ŷ�����             
	char uplmtprice                       [  7];	char _uplmtprice;                         //���Ѱ�               
	char high                             [  7];	char _high;                               //��                 
	char open                             [  7];	char _open;                               //�ð�                 
	char low                              [  7];	char _low;                                //����                 
	char dnlmtprice                       [  7];	char _dnlmtprice;                         //���Ѱ�               
} Tc4801OutBlock1;

typedef struct tagc4801OutBlock2    //�ֽļ�������ü��
{
	char dongsi                           [  1];	char _dongsi;                             //����ȣ������         
	char jeqprice                         [  7];	char _jeqprice;                           //����ü�ᰡ           
	char jeqsign                          [  1];	char _jeqsign;                            //����ü���ȣ         
	char jeqchange                        [  7];	char _jeqchange;                          //����ü������       
	char jeqchrate                        [  5];	char _jeqchrate;                          //����ü������       
} Tc4801OutBlock2;

typedef struct tagc4801
{
	Tc4801InBlock                     c4801inblock                          ;  //�⺻�Է� 
	Tc4801OutBlock                    c4801outblock                         ;  //�ֽļ���MASTER�⺻�ڷ� 
	Tc4801OutBlock1                   c4801outblock1                        ;  //�����ڻ� 
	Tc4801OutBlock2                   c4801outblock2                        ;  //�ֽļ�������ü�� 
} Tc4801;


typedef struct tagc4805InBlock    //�Էµ���Ÿ
{
	char fuitemz9                         [  9];	char _fuitemz9;                           //�Է��ڵ�             
} Tc4805InBlock;

typedef struct tagc4805OutUnder    //c4805OutUnder
{
	char shcode                           [  6];	char _shcode;                             //�����ڵ�             
	char hname                            [ 20];	char _hname;                              //�����               
	char price                            [  7];	char _price;                              //���簡               
	char sign                             [  1];	char _sign;                               //���Ϻ��ȣ           
	char change                           [  6];	char _change;                             //���Ϻ�               
} Tc4805OutUnder;

typedef struct tagc4805OutSMaster    //c4805OutSMaster
{
	char fuitemz8                         [  8];	char _fuitemz8;                           //�����ڵ�             
	char fuspcurr                         [  8];	char _fuspcurr;                           //����                 
	char fuspsign                         [  1];	char _fuspsign;                           //���Ϻ��ȣ           
	char fuspchange                       [  7];	char _fuspchange;                         //���Ϻ�               
	char fuspchrate                       [  5];	char _fuspchrate;                         //�����               
	char fuspopen                         [  7];	char _fuspopen;                           //�ð�                 
	char fusphigh                         [  7];	char _fusphigh;                           //��                 
	char fusplow                          [  7];	char _fusplow;                            //����                 
	char fuspvolall                       [  7];	char _fuspvolall;                         //�ŷ���               
	char fuspvalall                       [ 12];	char _fuspvalall;                         //�����ŷ����(õ��)   
	char fuspcurr1                        [  7];	char _fuspcurr1;                          //����������(�ٿ���)   
	char fuspcurr2                        [  7];	char _fuspcurr2;                          //����������(������)   
	char fuitem1                          [  8];	char _fuitem1;                            //�����ڵ�(�ٿ���)     
	char fuitem2                          [  8];	char _fuitem2;                            //�����ڵ�(������)     
	char fudynhprice                      [  7];	char _fudynhprice;                        //�ǽð����Ѱ�         
	char fudynlprice                      [  7];	char _fudynlprice;                        //�ǽð����Ѱ�         
	char fudynpriceflag                   [  1];	char _fudynpriceflag;                     //�����������ѿ���     
} Tc4805OutSMaster;

typedef struct tagc4805OutHoga3    //�ֽļ���ȣ��3
{
	char fuspfuitem                       [  8];	char _fuspfuitem;                         //�����ڵ�             
	char fusphname                        [ 50];	char _fusphname;                          //�ѱ۸�               
	char fusphotime                       [  8];	char _fusphotime;                         //ȣ���ð�             
	char offer                            [  8];	char _offer;                              //�ŵ��켱ȣ��         
	char bid                              [  8];	char _bid;                                //�ż��켱ȣ��         
	char offerjan                         [  6];	char _offerjan;                           //�ŵ��ܷ�             
	char bidjan                           [  6];	char _bidjan;                             //�ż��ܷ�             
	char S2offer                          [  8];	char _S2offer;                            //2���ŵ�ȣ��          
	char S2bid                            [  8];	char _S2bid;                              //2���ż�ȣ��          
	char S2offerjan                       [  6];	char _S2offerjan;                         //2���ŵ��ܷ�          
	char S2bidjan                         [  6];	char _S2bidjan;                           //2���ż��ܷ�          
	char S3offer                          [  8];	char _S3offer;                            //3���ŵ�ȣ��          
	char S3bid                            [  8];	char _S3bid;                              //3���ż�ȣ��          
	char S3offerjan                       [  6];	char _S3offerjan;                         //3���ŵ��ܷ�          
	char S3bidjan                         [  6];	char _S3bidjan;                           //3���ż��ܷ�          
	char S4offer                          [  8];	char _S4offer;                            //4���ŵ�ȣ��          
	char S4bid                            [  8];	char _S4bid;                              //4���ż�ȣ��          
	char S4offerjan                       [  6];	char _S4offerjan;                         //4���ŵ��ܷ�          
	char S4bidjan                         [  6];	char _S4bidjan;                           //4���ż��ܷ�          
	char S5offer                          [  8];	char _S5offer;                            //5���ŵ�ȣ��          
	char S5bid                            [  8];	char _S5bid;                              //5���ż�ȣ��          
	char S5offerjan                       [  6];	char _S5offerjan;                         //5���ŵ��ܷ�          
	char S5bidjan                         [  6];	char _S5bidjan;                           //5���ż��ܷ�          
	char S6offer                          [  8];	char _S6offer;                            //6���ŵ�ȣ��          
	char S6bid                            [  8];	char _S6bid;                              //6���ż�ȣ��          
	char S6offerjan                       [  6];	char _S6offerjan;                         //6���ŵ��ܷ�          
	char S6bidjan                         [  6];	char _S6bidjan;                           //6���ż��ܷ�          
	char S7offer                          [  8];	char _S7offer;                            //7���ŵ�ȣ��          
	char S7bid                            [  8];	char _S7bid;                              //7���ż�ȣ��          
	char S7offerjan                       [  6];	char _S7offerjan;                         //7���ŵ��ܷ�          
	char S7bidjan                         [  6];	char _S7bidjan;                           //7���ż��ܷ�          
	char S8offer                          [  8];	char _S8offer;                            //8���ŵ�ȣ��          
	char S8bid                            [  8];	char _S8bid;                              //8���ż�ȣ��          
	char S8offerjan                       [  6];	char _S8offerjan;                         //8���ŵ��ܷ�          
	char S8bidjan                         [  6];	char _S8bidjan;                           //8���ż��ܷ�          
	char S9offer                          [  8];	char _S9offer;                            //9���ŵ�ȣ��          
	char S9bid                            [  8];	char _S9bid;                              //9���ż�ȣ��          
	char S9offerjan                       [  6];	char _S9offerjan;                         //9���ŵ��ܷ�          
	char S9bidjan                         [  6];	char _S9bidjan;                           //9���ż��ܷ�          
	char S0offer                          [  8];	char _S0offer;                            //10���ŵ�ȣ��         
	char S0bid                            [  8];	char _S0bid;                              //10���ż�ȣ��         
	char S0offerjan                       [  6];	char _S0offerjan;                         //10���ŵ��ܷ�         
	char S0bidjan                         [  6];	char _S0bidjan;                           //10���ż��ܷ�         
	char offersu                          [  4];	char _offersu;                            //�ŵ��Ǽ�             
	char bidsu                            [  4];	char _bidsu;                              //�ż��Ǽ�             
	char S2offersu                        [  4];	char _S2offersu;                          //2���ŵ��Ǽ�          
	char S2bidsu                          [  4];	char _S2bidsu;                            //2���ż��Ǽ�          
	char S3offersu                        [  4];	char _S3offersu;                          //3���ŵ��Ǽ�          
	char S3bidsu                          [  4];	char _S3bidsu;                            //3���ż��Ǽ�          
	char S4offersu                        [  4];	char _S4offersu;                          //4���ŵ��Ǽ�          
	char S4bidsu                          [  4];	char _S4bidsu;                            //4���ż��Ǽ�          
	char S5offersu                        [  4];	char _S5offersu;                          //5���ŵ��Ǽ�          
	char S5bidsu                          [  4];	char _S5bidsu;                            //5���ż��Ǽ�          
	char S6offersu                        [  4];	char _S6offersu;                          //6���ŵ��Ǽ�          
	char S6bidsu                          [  4];	char _S6bidsu;                            //6���ż��Ǽ�          
	char S7offersu                        [  4];	char _S7offersu;                          //7���ŵ��Ǽ�          
	char S7bidsu                          [  4];	char _S7bidsu;                            //7���ż��Ǽ�          
	char S8offersu                        [  4];	char _S8offersu;                          //8���ŵ��Ǽ�          
	char S8bidsu                          [  4];	char _S8bidsu;                            //8���ż��Ǽ�          
	char S9offersu                        [  4];	char _S9offersu;                          //9���ŵ��Ǽ�          
	char S9bidsu                          [  4];	char _S9bidsu;                            //9���ż��Ǽ�          
	char S0offersu                        [  4];	char _S0offersu;                          //10���ŵ��Ǽ�         
	char S0bidsu                          [  4];	char _S0bidsu;                            //10���ż��Ǽ�         
	char tofferjan                        [  6];	char _tofferjan;                          //�Ѹŵ��ܷ�           
	char tobidjan                         [  6];	char _tobidjan;                           //�Ѹż��ܷ�           
	char toffersu                         [  5];	char _toffersu;                           //�Ѹŵ��Ǽ�           
	char tbidsu                           [  5];	char _tbidsu;                             //�Ѹż��Ǽ�           
	char undershcode                      [  6];	char _undershcode;                        //�����ڻ������ڵ�     
	char underhname                       [ 20];	char _underhname;                         //�����ڻ������       
	char eitem                            [  2];	char _eitem;                              //���ʴ���ֽ�         
	char lgcode                           [  9];	char _lgcode;                             //lgcode               
	char bp_jgubun                        [  1];	char _bp_jgubun;                          //BP���屸��           
} Tc4805OutHoga3;

typedef struct tagc4805OutSpread    //����SPREAD
{
	char thspread                         [  7];	char _thspread;                           //�̷н�������         
	char respread                         [  7];	char _respread;                           //������������         
	char fugrate1                         [  7];	char _fugrate1;                           //����                 
} Tc4805OutSpread;

typedef struct tagc4805
{
	Tc4805InBlock                     c4805inblock                          ;  //�Էµ���Ÿ 
	Tc4805OutUnder                    c4805outunder                         ;  //c4805OutUnder 
	Tc4805OutSMaster                  c4805outsmaster                       ;  //c4805OutSMaster 
	Tc4805OutHoga3                    c4805outhoga3                         ;  //�ֽļ���ȣ��3 
	Tc4805OutSpread                   c4805outspread                        ;  //����SPREAD 
} Tc4805;

typedef struct tags1701InBlock    //�⺻�Է�
{
	char code                             [  6];	char _code;                               //�����ڵ�             
} Ts1701InBlock;

typedef struct tags1701OutBlock    //���񸶽�Ÿ�⺻�ڷ�
{
	char code                             [  6];	char _code;                               //�����ڵ�             
	char hname                            [ 40];	char _hname;                              //�����               
	char price                            [  7];	char _price;                              //���簡               
	char sign                             [  1];	char _sign;                               //�����ȣ             
	char change                           [  6];	char _change;                             //�����               
	char chrate                           [  5];	char _chrate;                             //�����               
	char offer                            [  7];	char _offer;                              //�ŵ�ȣ��             
	char bid                              [  7];	char _bid;                                //�ż�ȣ��             
	char volume                           [  9];	char _volume;                             //�ŷ���               
	char volrate                          [  6];	char _volrate;                            //�ŷ�����             
	char value                            [  9];	char _value;                              //�ŷ����             
	char open                             [  7];	char _open;                               //�ð�                 
	char high                             [  7];	char _high;                               //��                 
	char low                              [  7];	char _low;                                //����                 
	char sale                             [  7];	char _sale;                               //���డ               
	char dnlmtprice                       [  7];	char _dnlmtprice;                         //���Ѱ�               
	char theoryprice                      [  7];	char _theoryprice;                        //�̷а�               
	char grate                            [  7];	char _grate;                              //������               
	char actprice                         [ 10];	char _actprice;                           //��簡               
	char listhprice                       [  7];	char _listhprice;                         //�������ְ�         
	char listhdate                        [  4];	char _listhdate;                          //�������ְ���       
	char listlprice                       [  7];	char _listlprice;                         //������������         
	char listldate                        [  4];	char _listldate;                          //��������������       
	char preprice                         [  7];	char _preprice;                           //��������             
	char hotime                           [  8];	char _hotime;                             //ȣ���ð�             
	char offho1                           [  7];	char _offho1;                             //�ŵ��ֿ켱ȣ��       
	char offho2                           [  7];	char _offho2;                             //�ŵ�����ȣ��         
	char offho3                           [  7];	char _offho3;                             //�ŵ�������ȣ��       
	char offho4                           [  7];	char _offho4;                             //�ŵ�4����ȣ��        
	char offho5                           [  7];	char _offho5;                             //�ŵ�5����ȣ��        
	char offho6                           [  7];	char _offho6;                             //�ŵ�6����ȣ��        
	char offho7                           [  7];	char _offho7;                             //�ŵ�7����ȣ��        
	char offho8                           [  7];	char _offho8;                             //�ŵ�8����ȣ��        
	char offho9                           [  7];	char _offho9;                             //�ŵ�9����ȣ��        
	char offho10                          [  7];	char _offho10;                            //�ŵ�10����ȣ��       
	char bidho1                           [  7];	char _bidho1;                             //�ż��ֿ켱ȣ��       
	char bidho2                           [  7];	char _bidho2;                             //�ż�����ȣ��         
	char bidho3                           [  7];	char _bidho3;                             //�ż�������ȣ��       
	char bidho4                           [  7];	char _bidho4;                             //�ż�4����ȣ��        
	char bidho5                           [  7];	char _bidho5;                             //�ż�5����ȣ��        
	char bidho6                           [  7];	char _bidho6;                             //�ż�6����ȣ��        
	char bidho7                           [  7];	char _bidho7;                             //�ż�7����ȣ��        
	char bidho8                           [  7];	char _bidho8;                             //�ż�8����ȣ��        
	char bidho9                           [  7];	char _bidho9;                             //�ż�9����ȣ��        
	char bidho10                          [  7];	char _bidho10;                            //�ż�10����ȣ��       
	char offremain1                       [  9];	char _offremain1;                         //�ŵ��ֿ켱�ܷ�       
	char offremain2                       [  9];	char _offremain2;                         //�ŵ������ܷ�         
	char offremain3                       [  9];	char _offremain3;                         //�ŵ��������ܷ�       
	char offremain4                       [  9];	char _offremain4;                         //�ŵ�4�����ܷ�        
	char offremain5                       [  9];	char _offremain5;                         //�ŵ�5�����ܷ�        
	char offremain6                       [  9];	char _offremain6;                         //�ŵ�6�����ܷ�        
	char offremain7                       [  9];	char _offremain7;                         //�ŵ�7�����ܷ�        
	char offremain8                       [  9];	char _offremain8;                         //�ŵ�8�����ܷ�        
	char offremain9                       [  9];	char _offremain9;                         //�ŵ�9�����ܷ�        
	char offremain10                      [  9];	char _offremain10;                        //�ŵ�10�����ܷ�       
	char bidremain1                       [  9];	char _bidremain1;                         //�ż��ֿ켱�ܷ�       
	char bidremain2                       [  9];	char _bidremain2;                         //�ż������ܷ�         
	char bidremain3                       [  9];	char _bidremain3;                         //�ż��������ܷ�       
	char bidremain4                       [  9];	char _bidremain4;                         //�ż�4�����ܷ�        
	char bidremain5                       [  9];	char _bidremain5;                         //�ż�5�����ܷ�        
	char bidremain6                       [  9];	char _bidremain6;                         //�ż�6�����ܷ�        
	char bidremain7                       [  9];	char _bidremain7;                         //�ż�7�����ܷ�        
	char bidremain8                       [  9];	char _bidremain8;                         //�ż�8�����ܷ�        
	char bidremain9                       [  9];	char _bidremain9;                         //�ż�9�����ܷ�        
	char bidremain10                      [  9];	char _bidremain10;                        //�ż�10�����ܷ�       
	char lpoffremain1                     [  9];	char _lpoffremain1;                       //LP�ŵ��ֿ켱�ܷ�     
	char lpoffremain2                     [  9];	char _lpoffremain2;                       //LP�ŵ������ܷ�       
	char lpoffremain3                     [  9];	char _lpoffremain3;                       //LP�ŵ��������ܷ�     
	char lpoffremain4                     [  9];	char _lpoffremain4;                       //LP�ŵ�4�����ܷ�      
	char lpoffremain5                     [  9];	char _lpoffremain5;                       //LP�ŵ�5�����ܷ�      
	char lpoffremain6                     [  9];	char _lpoffremain6;                       //LP�ŵ�6�����ܷ�      
	char lpoffremain7                     [  9];	char _lpoffremain7;                       //LP�ŵ�7�����ܷ�      
	char lpoffremain8                     [  9];	char _lpoffremain8;                       //LP�ŵ�8�����ܷ�      
	char lpoffremain9                     [  9];	char _lpoffremain9;                       //LP�ŵ�9�����ܷ�      
	char lpoffremain10                    [  9];	char _lpoffremain10;                      //LP�ŵ�10�����ܷ�     
	char lpbidremain1                     [  9];	char _lpbidremain1;                       //LP�ż��ֿ켱�ܷ�     
	char lpbidremain2                     [  9];	char _lpbidremain2;                       //LP�ż������ܷ�       
	char lpbidremain3                     [  9];	char _lpbidremain3;                       //LP�ż��������ܷ�     
	char lpbidremain4                     [  9];	char _lpbidremain4;                       //LP�ż�4�����ܷ�      
	char lpbidremain5                     [  9];	char _lpbidremain5;                       //LP�ż�5�����ܷ�      
	char lpbidremain6                     [  9];	char _lpbidremain6;                       //LP�ż�6�����ܷ�      
	char lpbidremain7                     [  9];	char _lpbidremain7;                       //LP�ż�7�����ܷ�      
	char lpbidremain8                     [  9];	char _lpbidremain8;                       //LP�ż�8�����ܷ�      
	char lpbidremain9                     [  9];	char _lpbidremain9;                       //LP�ż�9�����ܷ�      
	char lpbidremain10                    [  9];	char _lpbidremain10;                      //LP�ż�10�����ܷ�     
	char offtot                           [  9];	char _offtot;                             //�Ѹŵ��ܷ�           
	char bidtot                           [  9];	char _bidtot;                             //�Ѹż��ܷ�           
	char impv                             [ 10];	char _impv;                               //���纯����           
	char delta                            [  9];	char _delta;                              //��Ÿ����             
	char gamma                            [  9];	char _gamma;                              //��������             
	char vega                             [  9];	char _vega;                               //����������           
	char theta                            [  9];	char _theta;                              //��Ÿ�ð�             
	char rho                              [  9];	char _rho;                                //��������             
	char cdratio                          [  6];	char _cdratio;                            //������               
	char divideratio                      [  9];	char _divideratio;                        //��������           
	char jandatecnt                       [  4];	char _jandatecnt;                         //������               
	char elwsdate                         [  8];	char _elwsdate;                           //���Ⱓ������       
	char elwedate                         [  8];	char _elwedate;                           //���Ⱓ������       
	char lastdate                         [  8];	char _lastdate;                           //�����ŷ���           
	char balname                          [ 18];	char _balname;                            //������             
	char listing                          [  9];	char _listing;                            //�������             
	char rightgb                          [  4];	char _rightgb;                            //�Ǹ�����             /*��,ǲ,��Ÿ*/
	char righttype                        [  6];	char _righttype;                          //�Ǹ������         /*������,�̱���,��Ÿ*/
	char settletype                       [  9];	char _settletype;                         //�������             /*����,�ǹ�,����+�ǹ�*/
	char changerate                       [  8];	char _changerate;                         //��ȯ����             
	char rewardrate                       [  5];	char _rewardrate;                         //�ּ�������           
	char uppartrate                       [  5];	char _uppartrate;                         //���ݻ��������       
	char paydate                          [  8];	char _paydate;                            //����������           
	char lpjumun                          [  4];	char _lpjumun;                            //LP�ֹ����ɿ���       //�Ұ�,����
	char parity                           [  8];	char _parity;                             //�и�Ƽ               
	char gearingrate                      [  8];	char _gearingrate;                        //������           
	char profitrate                       [  8];	char _profitrate;                         //���ͺб���           
	char basepoint                        [  8];	char _basepoint;                          //�ں�������           
	char lp_name1                         [  6];	char _lp_name1;                           //LPȸ����1            
	char lp_name2                         [  6];	char _lp_name2;                           //LPȸ����2            
	char lp_name3                         [  6];	char _lp_name3;                           //LPȸ����3            
	char lp_name4                         [  6];	char _lp_name4;                           //LPȸ����4            
	char lp_name5                         [  6];	char _lp_name5;                           //LPȸ����5            
	char dongsi                           [  1];	char _dongsi;                             //����ȣ������         
	char eqprice                          [  7];	char _eqprice;                            //����ü�ᰡ           
	char eqsign                           [  1];	char _eqsign;                             //����ü���ȣ         
	char eqchange                         [  6];	char _eqchange;                           //����ü������       
	char eqchrate                         [  5];	char _eqchrate;                           //����ü������       
	char eqvol                            [  9];	char _eqvol;                              //����ü�����         
	char lphold                           [ 10];	char _lphold;                             //LP��������           
	char lprate                           [  5];	char _lprate;                             //LP������             
	char egearing                         [  8];	char _egearing;                           //E��              
	char fixpay                           [  8];	char _fixpay;                             //Ȯ�����޾�           
	char listdate                         [  8];	char _listdate;                           //������               
	char listhdatez8                      [  8];	char _listhdatez8;                        //�������ְ���       
	char listldatez8                      [  8];	char _listldatez8;                        //��������������       
	char intval                           [ 10];	char _intval;                             //���簡ġ             
	char leverage                         [  8];	char _leverage;                           //��������             
	char timeval                          [ 10];	char _timeval;                            //�ð���ġ             
	char gratio                           [  6];	char _gratio;                             //������               
	char profitpt                         [  8];	char _profitpt;                           //���ͺб���(����)     
	char payproxy                         [ 20];	char _payproxy;                           //���޴븮��           
	char standardopt                      [  2];	char _standardopt;                        //���񱸺�             /**01:ǥ��,03:��������**/
	char koprice                          [  6];	char _koprice;                            //�������ᰡ           
	char koappr                           [  5];	char _koappr;                             //KO���ٵ�             
	char expcode                          [ 12];	char _expcode;                            //Ȯ���ڵ�             
	char minpayment                       [  8];	char _minpayment;                         //�ּ����޾�           
	char stop                             [  1];	char _stop;                               //�ŷ���������         
	char gratio2                          [  8];	char _gratio2;                            //������2              
	char lpstop                           [  8];	char _lpstop;                             //LP������             
	char gonggb                           [  1];	char _gonggb;                             //�߰��������         
	char lp_impv                          [  5];	char _lp_impv;                            //LP���纯����         
	char r_intval                         [  8];	char _r_intval;                           //�ǽð��볻�簡ġ     
	char jandatecnt2                      [  4];	char _jandatecnt2;                        //������(������)       
	char profitpt2                        [ 10];	char _profitpt2;                          //���ͺб���(�Ҽ���)   
	char alertgb                          [  1];	char _alertgb;                            //�������Ǳ���         
} Ts1701OutBlock;

typedef struct tags1701OutBlock1    //�����ڻ�����, [�ݺ�]
{
	char code1                            [  6];	char _code1;                              //�����ڻ��ڵ�1        
	char hname1                           [ 20];	char _hname1;                             //�����ڻ��1          
	char price1                           [  7];	char _price1;                             //���簡1              
	char sign1                            [  1];	char _sign1;                              //�����ȣ1            
	char change1                          [  6];	char _change1;                            //�����1              
	char chrate1                          [  5];	char _chrate1;                            //�����1              
	char comrate1                         [  5];	char _comrate1;                           //������1              
	char pastv1                           [  5];	char _pastv1;                             //���ź�����1          
	char basegubun                        [  1];	char _basegubun;                          //�����ڻ���屸��     /*1:�ڽ���,2:�ڽ���*/
} Ts1701OutBlock1;

typedef struct tags1701OutBlock2    //�ŷ�������
{
	char tratimez5                        [  5];	char _tratimez5;                          //�ð�                 
	char off_tra1                         [  6];	char _off_tra1;                           //�ŵ��ŷ���1          
	char bid_tra1                         [  6];	char _bid_tra1;                           //�ż��ŷ���1          
	char offvolume1                       [  9];	char _offvolume1;                         //�ŵ��ŷ���1          
	char bidvolume1                       [  9];	char _bidvolume1;                         //�ż��ŷ���1          
	char off_tra2                         [  6];	char _off_tra2;                           //�ŵ��ŷ���2          
	char bid_tra2                         [  6];	char _bid_tra2;                           //�ż��ŷ���2          
	char offvolume2                       [  9];	char _offvolume2;                         //�ŵ��ŷ���2          
	char bidvolume2                       [  9];	char _bidvolume2;                         //�ż��ŷ���2          
	char off_tra3                         [  6];	char _off_tra3;                           //�ŵ��ŷ���3          
	char bid_tra3                         [  6];	char _bid_tra3;                           //�ż��ŷ���3          
	char offvolume3                       [  9];	char _offvolume3;                         //�ŵ��ŷ���3          
	char bidvolume3                       [  9];	char _bidvolume3;                         //�ż��ŷ���3          
	char off_tra4                         [  6];	char _off_tra4;                           //�ŵ��ŷ���4          
	char bid_tra4                         [  6];	char _bid_tra4;                           //�ż��ŷ���4          
	char offvolume4                       [  9];	char _offvolume4;                         //�ŵ��ŷ���4          
	char bidvolume4                       [  9];	char _bidvolume4;                         //�ż��ŷ���4          
	char off_tra5                         [  6];	char _off_tra5;                           //�ŵ��ŷ���5          
	char bid_tra5                         [  6];	char _bid_tra5;                           //�ż��ŷ���5          
	char offvolume5                       [  9];	char _offvolume5;                         //�ŵ��ŷ���5          
	char bidvolume5                       [  9];	char _bidvolume5;                         //�ż��ŷ���5          
	char offvolall                        [  9];	char _offvolall;                          //�ŵ��ܱ��ΰŷ���     
	char bidvolall                        [  9];	char _bidvolall;                          //�ż��ܱ��ΰŷ���     
	char alloffvol                        [  9];	char _alloffvol;                          //��ü�ŷ����ŵ���     
	char allbidvol                        [  9];	char _allbidvol;                          //��ü�ŷ����ż���     
} Ts1701OutBlock2;

typedef struct tags1701OutBlock3    //ELW�����ŷ����ڷ�, [�ݺ�]
{
	char chetime                          [  8];	char _chetime;                            //�ð�                 
	char price                            [  7];	char _price;                              //���簡               
	char sign                             [  1];	char _sign;                               //�����ȣ             
	char change                           [  6];	char _change;                             //�����               
	char offer                            [  7];	char _offer;                              //�ŵ�ȣ��             
	char bid                              [  7];	char _bid;                                //�ż�ȣ��             
	char volume                           [  9];	char _volume;                             //�ŷ���               
	char movvol                           [  9];	char _movvol;                             //������               
} Ts1701OutBlock3;

typedef struct tags1701OutBlock4    //K200�����ڻ�����
{
	char code6                            [  2];	char _code6;                              //�����ڻ��ڵ�6        
	char hname6                           [ 20];	char _hname6;                             //�����ڻ��6          
	char price6                           [  7];	char _price6;                             //���簡6              
	char sign6                            [  1];	char _sign6;                              //�����ȣ6            
	char change6                          [  6];	char _change6;                            //�����6              
	char chrate6                          [  5];	char _chrate6;                            //�����6              
	char comrate6                         [  5];	char _comrate6;                           //������6              
	char pastv6                           [  5];	char _pastv6;                             //���ź�����6          
} Ts1701OutBlock4;

typedef struct tags1701OutBlock5    //�ؿ����������ڻ�����
{
	char code7                            [  6];	char _code7;                              //�����ڻ��ڵ�7        
	char hname7                           [ 16];	char _hname7;                             //�����ڻ��7          
	char price7                           [  9];	char _price7;                             //���簡7              
	char sign7                            [  1];	char _sign7;                              //�����ȣ7            
	char change7                          [  9];	char _change7;                            //�����7              
	char chrate7                          [  5];	char _chrate7;                            //�����7              
	char time7                            [  4];	char _time7;                              //����Ÿ�ð�           
} Ts1701OutBlock5;

typedef struct tags1701
{
	Ts1701InBlock                     s1701inblock                          ;  //�⺻�Է� 
	Ts1701OutBlock                    s1701outblock                         ;  //���񸶽�Ÿ�⺻�ڷ� 
	Ts1701OutBlock1                   s1701outblock1                   [ 20];  //�����ڻ����� , [�ݺ�]
	Ts1701OutBlock2                   s1701outblock2                        ;  //�ŷ������� 
	Ts1701OutBlock3                   s1701outblock3                   [ 20];  //ELW�����ŷ����ڷ� , [�ݺ�]
	Ts1701OutBlock4                   s1701outblock4                        ;  //K200�����ڻ����� 
	Ts1701OutBlock5                   s1701outblock5                        ;  //�ؿ����������ڻ����� 
} Ts1701;

typedef struct tagp1003InBlock    //�Է�Data
{
	char formlang                         [  1];	char _formlang;                           //�ѿ�����             
	char gubun                            [  1];	char _gubun;                              //���ɱ���             /*f:KRX����,o:KRX�ɼ�,u:���μ���,p:���οɼ�*/
} Tp1003InBlock;

typedef struct tagp1003OutBlock    //�ڵ����Data, [�ݺ�]
{
	char codez8                           [  8];	char _codez8;                             //code                 
	char namez30                          [ 30];	char _namez30;                            //name                 
} Tp1003OutBlock;

typedef struct tagp1003
{
	Tp1003InBlock                     p1003inblock                          ;  //�Է�Data 
	Tp1003OutBlock                    p1003outblock                    [ 20];  //�ڵ����Data , [�ݺ�]
} Tp1003;

typedef struct tagc1151InBlock    //�⺻�Է�
{
	char formlang                         [  1];	char _formlang;                           //�ѿ�����             
	char code                             [  6];	char _code;                               //�����ڵ�             
} Tc1151InBlock;

typedef struct tagc1151OutBlock    //���񸶽�Ÿ�⺻�ڷ�
{
	char code                             [  6];	char _code;                               //�����ڵ�             
	char hname                            [ 13];	char _hname;                              //�����               
	char price                            [  7];	char _price;                              //���簡               
	char sign                             [  1];	char _sign;                               //�����ȣ             
	char change                           [  6];	char _change;                             //�����               
	char chrate                           [  5];	char _chrate;                             //�����               
	char offer                            [  7];	char _offer;                              //�ŵ�ȣ��             
	char bid                              [  7];	char _bid;                                //�ż�ȣ��             
	char volume                           [  9];	char _volume;                             //�ŷ���               
	char volrate                          [  6];	char _volrate;                            //�ŷ�����             
	char yurate                           [  5];	char _yurate;                             //������ȸ����         
	char value                            [  9];	char _value;                              //�ŷ����             
	char uplmtprice                       [  7];	char _uplmtprice;                         //���Ѱ�               
	char high                             [  7];	char _high;                               //���߰�             
	char open                             [  7];	char _open;                               //�ð�                 
	char opensign                         [  1];	char _opensign;                           //�ð�����ȣ         
	char openchange                       [  6];	char _openchange;                         //�ð��������       
	char low                              [  7];	char _low;                                //��������             
	char dnlmtprice                       [  7];	char _dnlmtprice;                         //���Ѱ�               
	char hotime                           [  8];	char _hotime;                             //ȣ���ð�             
	char offerho                          [  7];	char _offerho;                            //�ŵ��ֿ켱ȣ��       
	char P_offer                          [  7];	char _P_offer;                            //�ŵ�����ȣ��         
	char S_offer                          [  7];	char _S_offer;                            //�ŵ�������ȣ��       
	char S4_offer                         [  7];	char _S4_offer;                           //�ŵ�4����ȣ��        
	char S5_offer                         [  7];	char _S5_offer;                           //�ŵ�5����ȣ��        
	char S6_offer                         [  7];	char _S6_offer;                           //�ŵ�6����ȣ��        
	char S7_offer                         [  7];	char _S7_offer;                           //�ŵ�7����ȣ��        
	char S8_offer                         [  7];	char _S8_offer;                           //�ŵ�8����ȣ��        
	char S9_offer                         [  7];	char _S9_offer;                           //�ŵ�9����ȣ��        
	char S10_offer                        [  7];	char _S10_offer;                          //�ŵ�10����ȣ��       
	char bidho                            [  7];	char _bidho;                              //�ż��ֿ켱ȣ��       
	char P_bid                            [  7];	char _P_bid;                              //�ż�����ȣ��         
	char S_bid                            [  7];	char _S_bid;                              //�ż�������ȣ��       
	char S4_bid                           [  7];	char _S4_bid;                             //�ż�4����ȣ��        
	char S5_bid                           [  7];	char _S5_bid;                             //�ż�5����ȣ��        
	char S6_bid                           [  7];	char _S6_bid;                             //�ż�6����ȣ��        
	char S7_bid                           [  7];	char _S7_bid;                             //�ż�7����ȣ��        
	char S8_bid                           [  7];	char _S8_bid;                             //�ż�8����ȣ��        
	char S9_bid                           [  7];	char _S9_bid;                             //�ż�9����ȣ��        
	char S10_bid                          [  7];	char _S10_bid;                            //�ż�10����ȣ��       
	char offerrem                         [  9];	char _offerrem;                           //�ŵ��ֿ켱�ܷ�       
	char P_offerrem                       [  9];	char _P_offerrem;                         //�ŵ������ܷ�         
	char S_offerrem                       [  9];	char _S_offerrem;                         //�ŵ��������ܷ�       
	char S4_offerrem                      [  9];	char _S4_offerrem;                        //�ŵ�4�����ܷ�        
	char S5_offerrem                      [  9];	char _S5_offerrem;                        //�ŵ�5�����ܷ�        
	char S6_offerrem                      [  9];	char _S6_offerrem;                        //�ŵ�6�����ܷ�        
	char S7_offerrem                      [  9];	char _S7_offerrem;                        //�ŵ�7�����ܷ�        
	char S8_offerrem                      [  9];	char _S8_offerrem;                        //�ŵ�8�����ܷ�        
	char S9_offerrem                      [  9];	char _S9_offerrem;                        //�ŵ�9�����ܷ�        
	char S10_offerrem                     [  9];	char _S10_offerrem;                       //�ŵ�10�����ܷ�       
	char bidrem                           [  9];	char _bidrem;                             //�ż��ֿ켱�ܷ�       
	char P_bidrem                         [  9];	char _P_bidrem;                           //�ż������ܷ�         
	char S_bidrem                         [  9];	char _S_bidrem;                           //�ż��������ܷ�       
	char S4_bidrem                        [  9];	char _S4_bidrem;                          //�ż�4�����ܷ�        
	char S5_bidrem                        [  9];	char _S5_bidrem;                          //�ż�5�����ܷ�        
	char S6_bidrem                        [  9];	char _S6_bidrem;                          //�ż�6�����ܷ�        
	char S7_bidrem                        [  9];	char _S7_bidrem;                          //�ż�7�����ܷ�        
	char S8_bidrem                        [  9];	char _S8_bidrem;                          //�ż�8�����ܷ�        
	char S9_bidrem                        [  9];	char _S9_bidrem;                          //�ż�9�����ܷ�        
	char S10_bidrem                       [  9];	char _S10_bidrem;                         //�ż�10�����ܷ�       
	char T_offerrem                       [  9];	char _T_offerrem;                         //�Ѹŵ��ܷ�           
	char T_bidrem                         [  9];	char _T_bidrem;                           //�Ѹż��ܷ�           
	char O_offerrem                       [  9];	char _O_offerrem;                         //�ð��ܸŵ��ܷ�       
	char O_bidrem                         [  9];	char _O_bidrem;                           //�ð��ܸż��ܷ�       
	char pivot2upz7                       [  7];	char _pivot2upz7;                         //�Ǻ�2������          
	char pivot1upz7                       [  7];	char _pivot1upz7;                         //�Ǻ�1������          
	char pivotz7                          [  7];	char _pivotz7;                            //�Ǻ���               
	char pivot1dnz7                       [  7];	char _pivot1dnz7;                         //�Ǻ�1������          
	char pivot2dnz7                       [  7];	char _pivot2dnz7;                         //�Ǻ�2������          
	char sosokz6                          [  6];	char _sosokz6;                            //�ڽ����ڽ��ڱ���     
	char jisunamez18                      [ 18];	char _jisunamez18;                        //������               
	char capsizez6                        [  6];	char _capsizez6;                          //�ں��ݱԸ�           
	char output1z16                       [ 16];	char _output1z16;                         //����               
	char marcket1z16                      [ 16];	char _marcket1z16;                        //������ġ1            
	char marcket2z16                      [ 16];	char _marcket2z16;                        //������ġ2            
	char marcket3z16                      [ 16];	char _marcket3z16;                        //������ġ3            
	char marcket4z16                      [ 16];	char _marcket4z16;                        //������ġ4            
	char marcket5z16                      [ 16];	char _marcket5z16;                        //������ġ5            
	char marcket6z16                      [ 16];	char _marcket6z16;                        //������ġ6            
	char cbtext                           [  6];	char _cbtext;                             //CB����               
	char parvalue                         [  7];	char _parvalue;                           //�׸鰡               
	char prepricetitlez12                 [ 12];	char _prepricetitlez12;                   //��������Ÿ��Ʋ       
	char prepricez7                       [  7];	char _prepricez7;                         //��������             
	char subprice                         [  7];	char _subprice;                           //��밡               
	char gongpricez7                      [  7];	char _gongpricez7;                        //����               
	char high5                            [  7];	char _high5;                              //5�ϰ�              
	char low5                             [  7];	char _low5;                               //5������              
	char high20                           [  7];	char _high20;                             //20�ϰ�             
	char low20                            [  7];	char _low20;                              //20������             
	char yhigh                            [  7];	char _yhigh;                              //52���ְ�           
	char yhighdate                        [  4];	char _yhighdate;                          //52���ְ���         
	char ylow                             [  7];	char _ylow;                               //52��������           
	char ylowdate                         [  4];	char _ylowdate;                           //52����������         
	char movlistingz8                     [  8];	char _movlistingz8;                       //�����ֽļ�           
	char listing                          [ 12];	char _listing;                            //�����ֽļ�_õ��      
	char totpricez9                       [  9];	char _totpricez9;                         //�ð��Ѿ�             
	char tratimez5                        [  5];	char _tratimez5;                          //�ð�                 
	char off_tra1                         [  6];	char _off_tra1;                           //�ŵ��ŷ���1          
	char bid_tra1                         [  6];	char _bid_tra1;                           //�ż��ŷ���1          
	char N_offvolume1                     [  9];	char _N_offvolume1;                       //�ŵ��ŷ���1          
	char N_bidvolume1                     [  9];	char _N_bidvolume1;                       //�ż��ŷ���1          
	char off_tra2                         [  6];	char _off_tra2;                           //�ŵ��ŷ���2          
	char bid_tra2                         [  6];	char _bid_tra2;                           //�ż��ŷ���2          
	char N_offvolume2                     [  9];	char _N_offvolume2;                       //�ŵ��ŷ���2          
	char N_bidvolume2                     [  9];	char _N_bidvolume2;                       //�ż��ŷ���2          
	char off_tra3                         [  6];	char _off_tra3;                           //�ŵ��ŷ���3          
	char bid_tra3                         [  6];	char _bid_tra3;                           //�ż��ŷ���3          
	char N_offvolume3                     [  9];	char _N_offvolume3;                       //�ŵ��ŷ���3          
	char N_bidvolume3                     [  9];	char _N_bidvolume3;                       //�ż��ŷ���3          
	char off_tra4                         [  6];	char _off_tra4;                           //�ŵ��ŷ���4          
	char bid_tra4                         [  6];	char _bid_tra4;                           //�ż��ŷ���4          
	char N_offvolume4                     [  9];	char _N_offvolume4;                       //�ŵ��ŷ���4          
	char N_bidvolume4                     [  9];	char _N_bidvolume4;                       //�ż��ŷ���4          
	char off_tra5                         [  6];	char _off_tra5;                           //�ŵ��ŷ���5          
	char bid_tra5                         [  6];	char _bid_tra5;                           //�ż��ŷ���5          
	char N_offvolume5                     [  9];	char _N_offvolume5;                       //�ŵ��ŷ���5          
	char N_bidvolume5                     [  9];	char _N_bidvolume5;                       //�ż��ŷ���5          
	char N_offvolall                      [  9];	char _N_offvolall;                        //�ŵ��ܱ��ΰŷ���     
	char N_bidvolall                      [  9];	char _N_bidvolall;                        //�ż��ܱ��ΰŷ���     
	char fortimez6                        [  6];	char _fortimez6;                          //�ܱ��νð�           
	char forratez5                        [  5];	char _forratez5;                          //�ܱ���������         
	char settdatez4                       [  4];	char _settdatez4;                         //������               
	char cratez5                          [  5];	char _cratez5;                            //�ܰ����(%)          
	char yudatez4                         [  4];	char _yudatez4;                           //���������           
	char mudatez4                         [  4];	char _mudatez4;                           //���������           
	char yuratez5                         [  5];	char _yuratez5;                           //�����������         
	char muratez5                         [  5];	char _muratez5;                           //�����������         
	char listdatez8                       [  8];	char _listdatez8;                         //������               
	char listing2                         [ 12];	char _listing2;                           //�����ֽļ�_��        
	char N_alloffvol                      [  9];	char _N_alloffvol;                        //��ü�ŷ����ŵ���     
	char N_allbidvol                      [  9];	char _N_allbidvol;                        //��ü�ŷ����ż���     
} Tc1151OutBlock;

typedef struct tagc1151OutBlock2    //�����ŷ����ڷ�
{
	char time                             [  8];	char _time;                               //�ð�                 
	char price                            [  7];	char _price;                              //���簡               
	char sign                             [  1];	char _sign;                               //�����ȣ             
	char change                           [  6];	char _change;                             //�����               
	char offer                            [  7];	char _offer;                              //�ŵ�ȣ��             
	char bid                              [  7];	char _bid;                                //�ż�ȣ��             
	char movolume                         [  8];	char _movolume;                           //�����ŷ���           
	char volume                           [  9];	char _volume;                             //�ŷ���               
} Tc1151OutBlock2;

typedef struct tagc1151OutBlock3    //����ü��
{
	char dongsi                           [  1];	char _dongsi;                             //����ȣ������         
	char jeqprice                         [  7];	char _jeqprice;                           //����ü�ᰡ           
	char jeqsign                          [  1];	char _jeqsign;                            //����ü���ȣ         
	char jeqchange                        [  6];	char _jeqchange;                          //����ü������       
	char jeqchrate                        [  5];	char _jeqchrate;                          //����ü������       
	char jeqvol                           [  9];	char _jeqvol;                             //����ü�����         
} Tc1151OutBlock3;

typedef struct tagc1151OutBlock4    //ETF�ڷ�
{
	char bu12                             [  1];	char _bu12;                               //ETF����              
	char nav                              [  9];	char _nav;                                //����/����NAV         
	char nsign                            [  1];	char _nsign;                              //NAV�����ȣ          
	char nchange                          [  9];	char _nchange;                            //NAV�����            
	char prenav                           [  9];	char _prenav;                             //����NAVV             
	char grate                            [  9];	char _grate;                              //������               
	char gsign                            [  1];	char _gsign;                              //��������ȣ           
	char icuz18                           [ 18];	char _icuz18;                             //CU�����ݹ���(��)   
	char totjo                            [  4];	char _totjo;                              //���������           
	char totvalue                         [  7];	char _totvalue;                           //���ڻ��Ѿ�(���)     
	char terror                           [  9];	char _terror;                             //����������           
	char lpoffremain1                     [  9];	char _lpoffremain1;                       //LP�ŵ��ֿ켱�ܷ�     
	char lpoffremain2                     [  9];	char _lpoffremain2;                       //LP�ŵ������ܷ�       
	char lpoffremain3                     [  9];	char _lpoffremain3;                       //LP�ŵ��������ܷ�     
	char lpoffremain4                     [  9];	char _lpoffremain4;                       //LP�ŵ�4�����ܷ�      
	char lpoffremain5                     [  9];	char _lpoffremain5;                       //LP�ŵ�5�����ܷ�      
	char lpoffremain6                     [  9];	char _lpoffremain6;                       //LP�ŵ�6�����ܷ�      
	char lpoffremain7                     [  9];	char _lpoffremain7;                       //LP�ŵ�7�����ܷ�      
	char lpoffremain8                     [  9];	char _lpoffremain8;                       //LP�ŵ�8�����ܷ�      
	char lpoffremain9                     [  9];	char _lpoffremain9;                       //LP�ŵ�9�����ܷ�      
	char lpoffremain10                    [  9];	char _lpoffremain10;                      //LP�ŵ�10�����ܷ�     
	char lpbidremain1                     [  9];	char _lpbidremain1;                       //LP�ż��ֿ켱�ܷ�     
	char lpbidremain2                     [  9];	char _lpbidremain2;                       //LP�ż������ܷ�       
	char lpbidremain3                     [  9];	char _lpbidremain3;                       //LP�ż��������ܷ�     
	char lpbidremain4                     [  9];	char _lpbidremain4;                       //LP�ż�4�����ܷ�      
	char lpbidremain5                     [  9];	char _lpbidremain5;                       //LP�ż�5�����ܷ�      
	char lpbidremain6                     [  9];	char _lpbidremain6;                       //LP�ż�6�����ܷ�      
	char lpbidremain7                     [  9];	char _lpbidremain7;                       //LP�ż�7�����ܷ�      
	char lpbidremain8                     [  9];	char _lpbidremain8;                       //LP�ż�8�����ܷ�      
	char lpbidremain9                     [  9];	char _lpbidremain9;                       //LP�ż�9�����ܷ�      
	char lpbidremain10                    [  9];	char _lpbidremain10;                      //LP�ż�10�����ܷ�     
	char etf_copy_cd                      [  8];	char _etf_copy_cd;                        //ETF������������ڵ�  
	char etf_prod_cd                      [  6];	char _etf_prod_cd;                        //ETF��ǰ�����ڵ�      
} Tc1151OutBlock4;

typedef struct tagc1151OutBlock5    //���̽������ڷ�
{
	char jisucode                         [  2];	char _jisucode;                           //�����ڵ�             
	char sectorcode                       [  4];	char _sectorcode;                         //�����ڵ�             
	char jisuhnamez20                     [ 20];	char _jisuhnamez20;                       //������               
	char kp200jisu                        [  8];	char _kp200jisu;                          //����                 
	char kp200sign                        [  1];	char _kp200sign;                          //�����ȣ             
	char kp200change                      [  8];	char _kp200change;                        //�����               
	char ubjisu                           [ 10];	char _ubjisu;                             //ä������             
	char ubsign                           [  1];	char _ubsign;                             //ä�ǵ����ȣ         
	char ubchange                         [ 10];	char _ubchange;                           //ä�ǵ����           
	char symbol                           [ 12];	char _symbol;                             //�ؿ������ɺ�         
	char eupcode                          [  3];	char _eupcode;                            //��Ÿ�����ڵ�         
	char ubjiid                           [  6];	char _ubjiid;                             //ä�������ڵ�         
	char ubjiid2                          [  1];	char _ubjiid2;                            //ä�����������ڵ�     
} Tc1151OutBlock5;

typedef struct tagc1151
{
	Tc1151InBlock                     c1151inblock                          ;  //�⺻�Է� 
	Tc1151OutBlock                    c1151outblock                         ;  //���񸶽�Ÿ�⺻�ڷ� 
	Tc1151OutBlock2                   c1151outblock2                        ;  //�����ŷ����ڷ� 
	Tc1151OutBlock3                   c1151outblock3                        ;  //����ü�� 
	Tc1151OutBlock4                   c1151outblock4                        ;  //ETF�ڷ� 
	Tc1151OutBlock5                   c1151outblock5                        ;  //���̽������ڷ� 
} Tc1151;



typedef struct tagh1InBlock    //�Է�
{
	char code                             [  6];   //�����ڵ�             
} Th1InBlock;

typedef struct tagh1OutBlock    //���
{
	char code                             [  6];   //�����ڵ�             
	char hotime                           [  8];   //�ð�                 
	char offer                            [  7];   //�ŵ�ȣ��             
	char bid                              [  7];   //�ż�ȣ��             
	char offerrem                         [  9];   //�ŵ�ȣ���ܷ�         
	char bidrem                           [  9];   //�ż�ȣ���ܷ�         
	char P_offer                          [  7];   //���ŵ�ȣ��           
	char P_bid                            [  7];   //���ż�ȣ��           
	char P_offerrem                       [  9];   //���ŵ�ȣ���ܷ�       
	char P_bidrem                         [  9];   //���ż�ȣ���ܷ�       
	char S_offer                          [  7];   //�����ŵ�ȣ��         
	char S_bid                            [  7];   //�����ż�ȣ��         
	char S_offerrem                       [  9];   //�����ŵ�ȣ���ܷ�     
	char S_bidrem                         [  9];   //�����ż�ȣ���ܷ�     
	char S4_offer                         [  7];   //4���ŵ�ȣ��          
	char S4_bid                           [  7];   //4���ż�ȣ��          
	char S4_offerrem                      [  9];   //4���ŵ�ȣ���ܷ�      
	char S4_bidrem                        [  9];   //4���ż�ȣ���ܷ�      
	char S5_offer                         [  7];   //5���ŵ�ȣ��          
	char S5_bid                           [  7];   //5���ż�ȣ��          
	char S5_offerrem                      [  9];   //5���ŵ�ȣ���ܷ�      
	char S5_bidrem                        [  9];   //5���ż�ȣ���ܷ�      
	char T_offerrem                       [  9];   //�Ѹŵ�ȣ���ܷ�       
	char T_bidrem                         [  9];   //�Ѹż�ȣ���ܷ�       
	char S6_offer                         [  7];   //6���ŵ�ȣ��          
	char S6_bid                           [  7];   //6���ż�ȣ��          
	char S6_offerrem                      [  9];   //6���ŵ�ȣ���ܷ�      
	char S6_bidrem                        [  9];   //6���ż�ȣ���ܷ�      
	char S7_offer                         [  7];   //7���ŵ�ȣ��          
	char S7_bid                           [  7];   //7���ż�ȣ��          
	char S7_offerrem                      [  9];   //7���ŵ�ȣ���ܷ�      
	char S7_bidrem                        [  9];   //7���ż�ȣ���ܷ�      
	char S8_offer                         [  7];   //8���ŵ�ȣ��          
	char S8_bid                           [  7];   //8���ż�ȣ��          
	char S8_offerrem                      [  9];   //8���ŵ�ȣ���ܷ�      
	char S8_bidrem                        [  9];   //8���ż�ȣ���ܷ�      
	char S9_offer                         [  7];   //9���ŵ�ȣ��          
	char S9_bid                           [  7];   //9���ż�ȣ��          
	char S9_offerrem                      [  9];   //9���ŵ�ȣ���ܷ�      
	char S9_bidrem                        [  9];   //9���ż�ȣ���ܷ�      
	char S10_offer                        [  7];   //10���ŵ�ȣ��         
	char S10_bid                          [  7];   //10���ż�ȣ��         
	char S10_offerrem                     [  9];   //10���ŵ�ȣ���ܷ�     
	char S10_bidrem                       [  9];   //10���ż�ȣ���ܷ�     
	char volume                           [  9];   //�����ŷ���           
} Th1OutBlock;

typedef struct tagh1
{
	Th1InBlock                        h1inblock                             ;  //�Է� 
	Th1OutBlock                       h1outblock                            ;  //��� 
} Th1;

typedef struct tagk3InBlock    //�Է�
{
	char code                             [  6];   //�����ڵ�             
} Tk3InBlock;

typedef struct tagk3OutBlock    //���
{
	char code                             [  6];   //�����ڵ�             
	char time                             [  8];   //�ð�                 
	char offer                            [  7];   //�ŵ�ȣ��             
	char bid                              [  7];   //�ż�ȣ��             
	char offerrem                         [  9];   //�ŵ�ȣ���ܷ�         
	char bidrem                           [  9];   //�ż�ȣ���ܷ�         
	char P_offer                          [  7];   //���ŵ�ȣ��           
	char P_bid                            [  7];   //���ż�ȣ��           
	char P_offerrem                       [  9];   //���ŵ�ȣ���ܷ�       
	char P_bidrem                         [  9];   //���ż�ȣ���ܷ�       
	char S_offer                          [  7];   //�����ŵ�ȣ��         
	char S_bid                            [  7];   //�����ż�ȣ��         
	char S_offerrem                       [  9];   //�����ŵ�ȣ���ܷ�     
	char S_bidrem                         [  9];   //�����ż�ȣ���ܷ�     
	char S4_offer                         [  7];   //4���ŵ�ȣ��          
	char S4_bid                           [  7];   //4���ż�ȣ��          
	char S4_offerrem                      [  9];   //4���ŵ�ȣ���ܷ�      
	char S4_bidrem                        [  9];   //4���ż�ȣ���ܷ�      
	char S5_offer                         [  7];   //5���ŵ�ȣ��          
	char S5_bid                           [  7];   //5���ż�ȣ��          
	char S5_offerrem                      [  9];   //5���ŵ�ȣ���ܷ�      
	char S5_bidrem                        [  9];   //5���ż�ȣ���ܷ�      
	char T_offerrem                       [  9];   //�Ѹŵ�ȣ���ܷ�       
	char T_bidrem                         [  9];   //�Ѹż�ȣ���ܷ�       
	char S6_offer                         [  7];   //6���ŵ�ȣ��          
	char S6_bid                           [  7];   //6���ż�ȣ��          
	char S6_offerrem                      [  9];   //6���ŵ�ȣ���ܷ�      
	char S6_bidrem                        [  9];   //6���ż�ȣ���ܷ�      
	char S7_offer                         [  7];   //7���ŵ�ȣ��          
	char S7_bid                           [  7];   //7���ż�ȣ��          
	char S7_offerrem                      [  9];   //7���ŵ�ȣ���ܷ�      
	char S7_bidrem                        [  9];   //7���ż�ȣ���ܷ�      
	char S8_offer                         [  7];   //8���ŵ�ȣ��          
	char S8_bid                           [  7];   //8���ż�ȣ��          
	char S8_offerrem                      [  9];   //8���ŵ�ȣ���ܷ�      
	char S8_bidrem                        [  9];   //8���ż�ȣ���ܷ�      
	char S9_offer                         [  7];   //9���ŵ�ȣ��          
	char S9_bid                           [  7];   //9���ż�ȣ��          
	char S9_offerrem                      [  9];   //9���ŵ�ȣ���ܷ�      
	char S9_bidrem                        [  9];   //9���ż�ȣ���ܷ�      
	char S10_offer                        [  7];   //10���ŵ�ȣ��         
	char S10_bid                          [  7];   //10���ż�ȣ��         
	char S10_offerrem                     [  9];   //10���ŵ�ȣ���ܷ�     
	char S10_bidrem                       [  9];   //10���ż�ȣ���ܷ�     
	char volume                           [  9];   //�ŷ���               
} Tk3OutBlock;

typedef struct tagk3
{
	Tk3InBlock                        k3inblock                             ;  //�Է� 
	Tk3OutBlock                       k3outblock                            ;  //��� 
} Tk3;

typedef struct tagh2InBlock    //�Է�
{
	char code                             [  6];   //�����ڵ�             
} Th2InBlock;

typedef struct tagh2OutBlock    //���
{
	char code                             [  6];   //�����ڵ�             
	char hotime                           [  8];   //�ð�                 
	char O_offerrem                       [  9];   //�Ѹŵ�ȣ���ܷ�       
	char O_bidrem                         [  9];   //�Ѹż�ȣ���ܷ�       
} Th2OutBlock;

typedef struct tagh2
{
	Th2InBlock                        h2inblock                             ;  //�Է� 
	Th2OutBlock                       h2outblock                            ;  //��� 
} Th2;

typedef struct tagk4InBlock    //�Է�
{
	char code                             [  6];   //�����ڵ�             
} Tk4InBlock;

typedef struct tagk4OutBlock    //���
{
	char code                             [  6];   //�����ڵ�             
	char hotime                           [  8];   //�ð�                 
	char O_offerrem                       [  9];   //�Ѹŵ�ȣ���ܷ�       
	char O_bidrem                         [  9];   //�Ѹż�ȣ���ܷ�       
} Tk4OutBlock;

typedef struct tagk4
{
	Tk4InBlock                        k4inblock                             ;  //�Է� 
	Tk4OutBlock                       k4outblock                            ;  //��� 
} Tk4;

typedef struct tagh3InBlock    //�Է�
{
	char code                             [  6];   //�����ڵ�             
} Th3InBlock;

typedef struct tagh3OutBlock    //���
{
	char code                             [  6];   //�����ڵ�             
	char hotime                           [  8];   //�ð�                 
	char dongsi                           [  1];   //���ñ���             
	char jeqprice                         [  7];   //����ü�ᰡ           
	char jeqsign                          [  1];   //��������ȣ         
	char jeqchange                        [  6];   //��������           
	char jeqchrate                        [  5];   //��������           
	char jeqvol                           [  9];   //����ü�����         
	char offer                            [  7];   //�ŵ�ȣ��             
	char bid                              [  7];   //�ż�ȣ��             
	char offerrem                         [  9];   //�ŵ�ȣ���ܷ�         
	char bidrem                           [  9];   //�ż�ȣ���ܷ�         
} Th3OutBlock;

typedef struct tagh3
{
	Th3InBlock                        h3inblock                             ;  //�Է� 
	Th3OutBlock                       h3outblock                            ;  //��� 
} Th3;

typedef struct tagk5InBlock    //�Է�
{
	char code                             [  6];   //�����ڵ�             
} Tk5InBlock;

typedef struct tagk5OutBlock    //���
{
	char code                             [  6];   //�����ڵ�             
	char hotime                           [  8];   //�ð�                 
	char dongsi                           [  1];   //���ñ���             
	char jeqprice                         [  7];   //����ü�ᰡ           
	char jeqsign                          [  1];   //��������ȣ         
	char jeqchange                        [  6];   //��������           
	char jeqchrate                        [  5];   //��������           
	char jeqvol                           [  9];   //����ü�����         
	char offer                            [  7];   //�ŵ�ȣ��             
	char bid                              [  7];   //�ż�ȣ��             
	char offerrem                         [  9];   //�ŵ�ȣ���ܷ�         
	char bidrem                           [  9];   //�ż�ȣ���ܷ�         
} Tk5OutBlock;

typedef struct tagk5
{
	Tk5InBlock                        k5inblock                             ;  //�Է� 
	Tk5OutBlock                       k5outblock                            ;  //��� 
} Tk5;

typedef struct tagj8InBlock    //�Է�
{
	char code                             [  6];	char _code;                               //�����ڵ�             
} Tj8InBlock;

typedef struct tagj8OutBlock    //���
{
	char code                             [  6];	char _code;                               //�����ڵ�             
	char time                             [  8];	char _time;                               //�ð�                 
	char sign                             [  1];	char _sign;                               //�����ȣ             
	char change                           [  6];	char _change;                             //�����               
	char price                            [  7];	char _price;                              //���簡               
	char chrate                           [  5];	char _chrate;                             //�����               
	char high                             [  7];	char _high;                               //��                 
	char low                              [  7];	char _low;                                //����                 
	char offer                            [  7];	char _offer;                              //�ŵ�ȣ��             
	char bid                              [  7];	char _bid;                                //�ż�ȣ��             
	char volume                           [  9];	char _volume;                             //�ŷ���               
	char volrate                          [  6];	char _volrate;                            //�ŷ������Ϻ�         
	char movolume                         [  8];	char _movolume;                           //�����ŷ���           
	char value                            [  9];	char _value;                              //�ŷ����             
	char open                             [  7];	char _open;                               //�ð�                 
	char avgprice                         [  7];	char _avgprice;                           //������հ�           
	char janggubun                        [  1];	char _janggubun;                          //�屸��               
} Tj8OutBlock;

typedef struct tagj8
{
	Tj8InBlock                        j8inblock                             ;  //�Է� 
	Tj8OutBlock                       j8outblock                            ;  //��� 
} Tj8;

typedef struct tagk8InBlock    //�Է�
{
	char code                             [  6];	char _code;                               //�����ڵ�             
} Tk8InBlock;

typedef struct tagk8OutBlock    //���
{
	char code                             [  6];	char _code;                               //�����ڵ�             
	char time                             [  8];	char _time;                               //�ð�                 
	char price                            [  7];	char _price;                              //���簡               
	char sign                             [  1];	char _sign;                               //�����ȣ             
	char change                           [  6];	char _change;                             //�����               
	char chrate                           [  5];	char _chrate;                             //�����               
	char high                             [  7];	char _high;                               //��                 
	char low                              [  7];	char _low;                                //����                 
	char offer                            [  7];	char _offer;                              //�ŵ�ȣ��             
	char bid                              [  7];	char _bid;                                //�ż�ȣ��             
	char volume                           [  9];	char _volume;                             //�ŷ���               
	char volrate                          [  6];	char _volrate;                            //�ŷ������Ϻ�         
	char movolume                         [  8];	char _movolume;                           //�����ŷ���           
	char value                            [  9];	char _value;                              //�ŷ����             
	char open                             [  7];	char _open;                               //�ð�                 
	char avgprice                         [  7];	char _avgprice;                           //������հ�           
	char janggubun                        [  1];	char _janggubun;                          //�屸��               
} Tk8OutBlock;

typedef struct tagk8
{
	Tk8InBlock                        k8inblock                             ;  //�Է� 
	Tk8OutBlock                       k8outblock                            ;  //��� 
} Tk8;

typedef struct tagf1InBlock    //�Է�
{
	char fuitem                           [  4];   //�����ڵ�             
} Tf1InBlock;

typedef struct tagf1OutBlock    //���
{
	char fuitem                           [  4];   //�����ڵ�             
	char fuhotime                         [  8];   //�ð�                 
	char fuoffer                          [  5];   //�ŵ��켱ȣ��         
	char fubid                            [  5];   //�ż��켱ȣ��         
	char fuofferjan                       [  6];   //�ŵ��켱�ܷ�         
	char fubidjan                         [  6];   //�ż��켱�ܷ�         
	char fujoffer                         [  5];   //�����ŵ�ȣ��         
	char fujbid                           [  5];   //�����ż�ȣ��         
	char fujofferjan                      [  6];   //�����ŵ��ܷ�         
	char fujbidjan                        [  6];   //�����ż��ܷ�         
	char fujjoffer                        [  5];   //�������ŵ�ȣ��       
	char fujjbid                          [  5];   //�������ż�ȣ��       
	char fujjofferjan                     [  6];   //�������ŵ��ܷ�       
	char fujjbidjan                       [  6];   //�������ż��ܷ�       
	char futofferjan                      [  6];   //�Ѹŵ��ܷ�           
	char futbidjan                        [  6];   //�Ѹż��ܷ�           
	char fuj4offer                        [  5];   //4�����ŵ�ȣ��        
	char fuj4bid                          [  5];   //4�����ż�ȣ��        
	char fuj4offerjan                     [  6];   //4�����ŵ��ܷ�        
	char fuj4bidjan                       [  6];   //4�����ż��ܷ�        
	char fuj5offer                        [  5];   //5�����ŵ�ȣ��        
	char fuj5bid                          [  5];   //5�����ż�ȣ��        
	char fuj5offerjan                     [  6];   //5�����ŵ��ܷ�        
	char fuj5bidjan                       [  6];   //5�����ż��ܷ�        
	char fuoffersu                        [  4];   //�켱�ŵ��Ǽ�         
	char fujoffersu                       [  4];   //�����ŵ��Ǽ�         
	char fujjoffersu                      [  4];   //�������ŵ��Ǽ�       
	char fuj4offersu                      [  4];   //4�����ŵ��Ǽ�        
	char fuj5offersu                      [  4];   //5�����ŵ��Ǽ�        
	char futoffersu                       [  5];   //�Ѹŵ��Ǽ�           
	char fubidsu                          [  4];   //�켱�ż��Ǽ�         
	char fujbidsu                         [  4];   //�����ż��Ǽ�         
	char fujjbidsu                        [  4];   //�������ż��Ǽ�       
	char fuj4bidsu                        [  4];   //4�����ż��Ǽ�        
	char fuj5bidsu                        [  4];   //5�����ż��Ǽ�        
	char futbidsu                         [  5];   //�Ѹż��Ǽ�           
} Tf1OutBlock;

typedef struct tagf1
{
	Tf1InBlock                        f1inblock                             ;  //�Է� 
	Tf1OutBlock                       f1outblock                            ;  //��� 
} Tf1;

typedef struct tagf3InBlock    //�Է�
{
	char fuitem                           [  4];   //�����ڵ�             
} Tf3InBlock;

typedef struct tagf3OutBlock    //���
{
	char fuitem                           [  4];   //�����ڵ�             
	char futheoryprice                    [  5];   //�����̷а�           
	char futheorytime                     [  8];   //�̷а��ð�           
	char fugrate                          [  5];   //������               
	char fugratio                         [  5];   //������               
} Tf3OutBlock;

typedef struct tagf3
{
	Tf3InBlock                        f3inblock                             ;  //�Է� 
	Tf3OutBlock                       f3outblock                            ;  //��� 
} Tf3;

typedef struct tagf4InBlock    //�Է�
{
	char fuitem                           [  4];   //�����ڵ�             
} Tf4InBlock;

typedef struct tagf4OutBlock    //���
{
	char fuitem                           [  4];   //�����ڵ�             
	char fuchetime                        [  8];   //ü��ð�             
	char fuopenyak                        [  7];   //�̰�����������       
	char fupreopenyak                     [  7];   //���Ϲ̰�����������   
} Tf4OutBlock;

typedef struct tagf4
{
	Tf4InBlock                        f4inblock                             ;  //�Է� 
	Tf4OutBlock                       f4outblock                            ;  //��� 
} Tf4;

typedef struct tagf8InBlock    //�Է�
{
	char fuitem                           [  4];	char _fuitem;                             //�����ڵ�             
} Tf8InBlock;

typedef struct tagf8OutBlock    //���
{
	char fuitem                           [  4];	char _fuitem;                             //�����ڵ�             
	char fuchetime                        [  8];	char _fuchetime;                          //�ð�                 
	char fusign                           [  1];	char _fusign;                             //�����ȣ             
	char fuchange                         [  5];	char _fuchange;                           //�����               
	char fucurr                           [  5];	char _fucurr;                             //���簡               
	char fuhigh                           [  5];	char _fuhigh;                             //��                 
	char fulow                            [  5];	char _fulow;                              //����                 
	char fuvol                            [  6];	char _fuvol;                              //ü�����             
	char fuvolall                         [  7];	char _fuvolall;                           //����ü�����         
	char fuvalall                         [ 12];	char _fuvalall;                           //�����ŷ����         
	char fuopenyak                        [  7];	char _fuopenyak;                          //�̰�����������       
	char fuoffer                          [  5];	char _fuoffer;                            //�켱�ŵ�ȣ��         
	char fubid                            [  5];	char _fubid;                              //�켱�ż�ȣ��         
	char fuofferjan                       [  6];	char _fuofferjan;                         //�켱�ŵ��ܷ�         
	char fubidjan                         [  6];	char _fubidjan;                           //�켱�ż��ܷ�         
	char futofferjan                      [  6];	char _futofferjan;                        //�Ѹŵ��ܷ�           
	char futbidjan                        [  6];	char _futbidjan;                          //�Ѹż��ܷ�           
	char fuoffersu                        [  4];	char _fuoffersu;                          //�켱�ŵ��Ǽ�         
	char fubidsu                          [  4];	char _fubidsu;                            //�켱�ż��Ǽ�         
	char futoffersu                       [  5];	char _futoffersu;                         //�Ѹŵ��Ǽ�           
	char futbidsu                         [  5];	char _futbidsu;                           //�Ѹż��Ǽ�           
	char fuchrate                         [  5];	char _fuchrate;                           //�����               
	char fubasis                          [  5];	char _fubasis;                            //���̽ý�             
	char fugrate                          [  5];	char _fugrate;                            //������               
	char fugratio                         [  5];	char _fugratio;                           //������               
	char fupreopenyak                     [  7];	char _fupreopenyak;                       //�̰�����������       
	char fuspvolall                       [  7];	char _fuspvolall;                         //�����������         
	char fuopen                           [  5];	char _fuopen;                             //�ð�                 
	char bulkvol                          [  7];	char _bulkvol;                            //���Ǵ뷮����ü����� 
} Tf8OutBlock;

typedef struct tagf8
{
	Tf8InBlock                        f8inblock                             ;  //�Է� 
	Tf8OutBlock                       f8outblock                            ;  //��� 
} Tf8;

typedef struct tagq1InBlock    //�Է�
{
	char fuspcode                         [  8];   //�����ڵ�             
} Tq1InBlock;

typedef struct tagq1OutBlock    //���
{
	char fuspcode                         [  8];   //�����ڵ�             
	char fusphotime                       [  8];   //�ð�                 
	char fuspoffer                        [  6];   //�켱�ŵ�ȣ��         
	char fuspbid                          [  6];   //�켱�ż�ȣ��         
	char fuspofferjan                     [  6];   //�켱�ŵ��ܷ�         
	char fuspbidjan                       [  6];   //�켱�ż��ܷ�         
	char fuspjoffer                       [  6];   //�����ŵ�ȣ��         
	char fuspjbid                         [  6];   //�����ż�ȣ��         
	char fuspjofferjan                    [  6];   //�����ŵ��ܷ�         
	char fuspjbidjan                      [  6];   //�����ż��ܷ�         
	char fuspjjoffer                      [  6];   //�������ŵ�ȣ��       
	char fuspjjbid                        [  6];   //�������ż�ȣ��       
	char fuspjjofferjan                   [  6];   //�������ŵ��ܷ�       
	char fuspjjbidjan                     [  6];   //�������ż��ܷ�       
	char fuspj4offer                      [  6];   //4�����ŵ�ȣ��        
	char fuspj4bid                        [  6];   //4�����ż�ȣ��        
	char fuspj4offerjan                   [  6];   //4�����ŵ��ܷ�        
	char fuspj4bidjan                     [  6];   //4�����ż��ܷ�        
	char fuspj5offer                      [  6];   //5�����ŵ�ȣ��        
	char fuspj5bid                        [  6];   //5�����ż�ȣ��        
	char fuspj5offerjan                   [  6];   //5�����ŵ��ܷ�        
	char fuspj5bidjan                     [  6];   //5�����ż��ܷ�        
	char fusptofferjan                    [  6];   //�Ѹŵ��ܷ�           
	char fusptbidjan                      [  6];   //�Ѹż��ܷ�           
	char fuspoffersu                      [  4];   //�켱�ŵ��Ǽ�         
	char fuspjoffersu                     [  4];   //�����ŵ��Ǽ�         
	char fuspjjoffersu                    [  4];   //�������ŵ��Ǽ�       
	char fuspj4offersu                    [  4];   //4�����ŵ��Ǽ�        
	char fuspj5offersu                    [  4];   //5�����ŵ��Ǽ�        
	char fusptoffersu                     [  5];   //�Ѹŵ��Ǽ�           
	char fuspbidsu                        [  4];   //�켱�ż��Ǽ�         
	char fuspjbidsu                       [  4];   //�����ż��Ǽ�         
	char fuspjjbidsu                      [  4];   //�������ż��Ǽ�       
	char fuspj4bidsu                      [  4];   //4�����ż��Ǽ�        
	char fuspj5bidsu                      [  4];   //5�����ż��Ǽ�        
	char fusptbidsu                       [  5];   //�Ѹż��Ǽ�           
} Tq1OutBlock;

typedef struct tagq1
{
	Tq1InBlock                        q1inblock                             ;  //�Է� 
	Tq1OutBlock                       q1outblock                            ;  //��� 
} Tq1;

typedef struct tagq2InBlock    //�Է�
{
	char fuspcode                         [  8];	char _fuspcode;                           //�����ڵ�             
} Tq2InBlock;

typedef struct tagq2OutBlock    //���
{
	char fuspcode                         [  8];	char _fuspcode;                           //�����ڵ�             
	char fusphotime                       [  8];	char _fusphotime;                         //�ð�                 
	char fuspjgubun                       [  8];	char _fuspjgubun;                         //����               
	char fuspsign                         [  1];	char _fuspsign;                           //���Ϻ�ȣ             
	char fuspchange                       [  5];	char _fuspchange;                         //���ϴ��             
	char fuspcurr                         [  6];	char _fuspcurr;                           //���簡               
	char fuspcurr1                        [  5];	char _fuspcurr1;                          //����������(�ٿ�)     
	char fuspcurr2                        [  5];	char _fuspcurr2;                          //����������(����)     
	char fuspopen                         [  6];	char _fuspopen;                           //�ð�                 
	char fusphigh                         [  6];	char _fusphigh;                           //��                 
	char fusplow                          [  6];	char _fusplow;                            //����                 
	char fuspvol                          [  6];	char _fuspvol;                            //ü�����             
	char fuspvolall                       [  7];	char _fuspvolall;                         //����ü�����         
	char fuspvalall                       [ 12];	char _fuspvalall;                         //�����ŷ����         
	char fuspchrate                       [  5];	char _fuspchrate;                         //�����               
	char fuspbp_jgubun                    [  1];	char _fuspbp_jgubun;                      //BP���屸��           
	char fuspoffer                        [  6];	char _fuspoffer;                          //�켱�ŵ�ȣ��         
	char fuspbid                          [  6];	char _fuspbid;                            //�켱�ż�ȣ��         
} Tq2OutBlock;

typedef struct tagq2
{
	Tq2InBlock                        q2inblock                             ;  //�Է� 
	Tq2OutBlock                       q2outblock                            ;  //��� 
} Tq2;

typedef struct tago1InBlock    //�Է�
{
	char opitem                           [  8];   //�ڵ�                 
} To1InBlock;

typedef struct tago1OutBlock    //���
{
	char opitem                           [  8];   //�ڵ�                 
	char ophotime                         [  8];   //�ð�                 
	char opoffer                          [  5];   //�ŵ��켱ȣ��         
	char opbid                            [  5];   //�ż��켱ȣ��         
	char opofferjan                       [  7];   //�ŵ��켱�ܷ�         
	char opbidjan                         [  7];   //�ż��켱�ܷ�         
	char opjoffer                         [  5];   //�����ŵ�ȣ��         
	char opjbid                           [  5];   //�����ż�ȣ��         
	char opjofferjan                      [  7];   //�����ŵ��ܷ�         
	char opjbidjan                        [  7];   //�����ż��ܷ�         
	char opjjoffer                        [  5];   //�������ŵ�ȣ��       
	char opjjbid                          [  5];   //�������ż�ȣ��       
	char opjjofferjan                     [  7];   //�������ŵ��ܷ�       
	char opjjbidjan                       [  7];   //�������ż��ܷ�       
	char optofferjan                      [  7];   //�Ѹŵ��ܷ�           
	char optbidjan                        [  7];   //�Ѹż��ܷ�           
	char opj4offer                        [  5];   //4�����ŵ�ȣ��        
	char opj4bid                          [  5];   //4�����ż�ȣ��        
	char opj4offerjan                     [  7];   //4�����ŵ��ܷ�        
	char opj4bidjan                       [  7];   //4�����ż��ܷ�        
	char opj5offer                        [  5];   //5�����ŵ�ȣ��        
	char opj5bid                          [  5];   //5�����ż�ȣ��        
	char opj5offerjan                     [  7];   //5�����ŵ��ܷ�        
	char opj5bidjan                       [  7];   //5�����ż��ܷ�        
	char opoffersu                        [  4];   //�켱�ŵ��Ǽ�         
	char opjoffersu                       [  4];   //�����ŵ��Ǽ�         
	char opjjoffersu                      [  4];   //�������ŵ��Ǽ�       
	char opj4offersu                      [  4];   //4�����ŵ��Ǽ�        
	char opj5offersu                      [  4];   //5�����ŵ��Ǽ�        
	char optoffersu                       [  5];   //�Ѹŵ��Ǽ�           
	char opbidsu                          [  4];   //�켱�ż��Ǽ�         
	char opjbidsu                         [  4];   //�����ż��Ǽ�         
	char opjjbidsu                        [  4];   //�������ż��Ǽ�       
	char opj4bidsu                        [  4];   //4�����ż��Ǽ�        
	char opj5bidsu                        [  4];   //5�����ż��Ǽ�        
	char optbidsu                         [  5];   //�Ѹż��Ǽ�           
} To1OutBlock;

typedef struct tago1
{
	To1InBlock                        o1inblock                             ;  //�Է� 
	To1OutBlock                       o1outblock                            ;  //��� 
} To1;

typedef struct tago2InBlock    //�Է�
{
	char opitem                           [  8];	char _opitem;                             //�����ڵ�             
} To2InBlock;

typedef struct tago2OutBlock    //���
{
	char opitem                           [  8];	char _opitem;                             //�����ڵ�             
	char opchetime                        [  8];	char _opchetime;                          //�ð�                 
	char opjgubun                         [  8];	char _opjgubun;                           //����               
	char opsign                           [  1];	char _opsign;                             //�����ȣ             
	char opchange                         [  5];	char _opchange;                           //�����               
	char opcurr                           [  5];	char _opcurr;                             //���簡               
	char opopen                           [  5];	char _opopen;                             //�ð�                 
	char ophigh                           [  5];	char _ophigh;                             //��                 
	char oplow                            [  5];	char _oplow;                              //����                 
	char opvol                            [  6];	char _opvol;                              //ü�����             
	char opvolallz8                       [  8];	char _opvolallz8;                         //����ü�����         
	char opvalall                         [ 12];	char _opvalall;                           //�����ŷ����         
	char opopenyak                        [  7];	char _opopenyak;                          //�̰�����������       
	char opoffer                          [  5];	char _opoffer;                            //�켱�ŵ�ȣ��         
	char opbid                            [  5];	char _opbid;                              //�켱�ż�ȣ��         
	char opofferjan                       [  7];	char _opofferjan;                         //�켱�ŵ��ܷ�         
	char opbidjan                         [  7];	char _opbidjan;                           //�켱�ż��ܷ�         
	char opjoffer                         [  5];	char _opjoffer;                           //�����ŵ�ȣ��         
	char opjbid                           [  5];	char _opjbid;                             //�����ż�ȣ��         
	char opjofferjan                      [  7];	char _opjofferjan;                        //�����ŵ��ܷ�         
	char opjbidjan                        [  7];	char _opjbidjan;                          //�����ż��ܷ�         
	char opjjoffer                        [  5];	char _opjjoffer;                          //�������ŵ�ȣ��       
	char opjjbid                          [  5];	char _opjjbid;                            //�������ż�ȣ��       
	char opjjofferjan                     [  7];	char _opjjofferjan;                       //�������ŵ��ܷ�       
	char opjjbidjan                       [  7];	char _opjjbidjan;                         //�������ż��ܷ�       
	char optofferjan                      [  7];	char _optofferjan;                        //�Ѹŵ��ܷ�           
	char optbidjan                        [  7];	char _optbidjan;                          //�Ѹż��ܷ�           
	char opj4offer                        [  5];	char _opj4offer;                          //4�����ŵ�ȣ��        
	char opj4bid                          [  5];	char _opj4bid;                            //4�����ż�ȣ��        
	char opj4offerjan                     [  7];	char _opj4offerjan;                       //4�����ŵ��ܷ�        
	char opj4bidjan                       [  7];	char _opj4bidjan;                         //4�����ż��ܷ�        
	char opj5offer                        [  5];	char _opj5offer;                          //5�����ŵ�ȣ��        
	char opj5bid                          [  5];	char _opj5bid;                            //5�����ż�ȣ��        
	char opj5offerjan                     [  7];	char _opj5offerjan;                       //5�����ŵ��ܷ�        
	char opj5bidjan                       [  7];	char _opj5bidjan;                         //5�����ż��ܷ�        
	char opoffersu                        [  4];	char _opoffersu;                          //�켱�ŵ��Ǽ�         
	char opjoffersu                       [  4];	char _opjoffersu;                         //�����ŵ��Ǽ�         
	char opjjoffersu                      [  4];	char _opjjoffersu;                        //�������ŵ��Ǽ�       
	char opj4offersu                      [  4];	char _opj4offersu;                        //4�����ŵ��Ǽ�        
	char opj5offersu                      [  4];	char _opj5offersu;                        //5�����ŵ��Ǽ�        
	char optoffersu                       [  5];	char _optoffersu;                         //�Ѹŵ��Ǽ�           
	char opbidsu                          [  4];	char _opbidsu;                            //�켱�ż��Ǽ�         
	char opjbidsu                         [  4];	char _opjbidsu;                           //�����ż��Ǽ�         
	char opjjbidsu                        [  4];	char _opjjbidsu;                          //�������ż��Ǽ�       
	char opj4bidsu                        [  4];	char _opj4bidsu;                          //4�����ż��Ǽ�        
	char opj5bidsu                        [  4];	char _opj5bidsu;                          //5�����ż��Ǽ�        
	char optbidsu                         [  5];	char _optbidsu;                           //�Ѹż��Ǽ�           
	char opchrate                         [  5];	char _opchrate;                           //�����               
	char opgrate                          [  5];	char _opgrate;                            //������               
	char opgratio                         [  5];	char _opgratio;                           //������               
	char oppreopenyak                     [  7];	char _oppreopenyak;                       //�̰�����������       
	char opbp_jgubun                      [  1];	char _opbp_jgubun;                        //BP���屸��           
	char bulkvolz8                        [  8];	char _bulkvolz8;                          //���Ǵ뷮����ü����� 
} To2OutBlock;

typedef struct tago2
{
	To2InBlock                        o2inblock                             ;  //�Է� 
	To2OutBlock                       o2outblock                            ;  //��� 
} To2;

typedef struct tago3InBlock    //�Է�
{
	char opitem                           [  8];   //�����ڵ�             
} To3InBlock;

typedef struct tago3OutBlock    //���
{
	char opitem                           [  8];   //�����ڵ�             
	char optheorytime                     [  8];   //�̷а��ð�           
	char optheoryprice                    [  5];   //�ɼ��̷а�           
	char opimpv                           [  5];   //���纯����           
	char opdelta                          [  8];   //��ȣ+��Ÿ            
	char opgmma                           [  8];   //��ȣ+����            
	char opvega                           [  8];   //��ȣ+����            
	char optheta                          [  8];   //��ȣ+��Ÿ            
	char oprho                            [  8];   //��ȣ+��              
	char opgrate                          [  5];   //������               
	char opgratio                         [  5];   //������               
} To3OutBlock;

typedef struct tago3
{
	To3InBlock                        o3inblock                             ;  //�Է� 
	To3OutBlock                       o3outblock                            ;  //��� 
} To3;

typedef struct tago4InBlock    //�Է�
{
	char opitem                           [  8];   //�����ڵ�             
} To4InBlock;

typedef struct tago4OutBlock    //���
{
	char opitem                           [  8];   //�����ڵ�             
	char opchetime                        [  8];   //ü��ð�             
	char opopenyak                        [  7];   //�̰�����������       
	char oppreopenyak                     [  7];   //���Ϲ̰�����������   
} To4OutBlock;

typedef struct tago4
{
	To4InBlock                        o4inblock                             ;  //�Է� 
	To4OutBlock                       o4outblock                            ;  //��� 
} To4;

typedef struct tagvHInBlock    //�Է�
{
	char fuitem                           [  8];   //�����ڵ�             
} TvHInBlock;

typedef struct tagvHOutBlock    //���
{
	char fuitem                           [  8];   //�����ڵ�             
	char futime                           [  8];   //�ð� HH:MM:SS        
	char offer                            [  7];   //�ŵ�ȣ��             
	char bid                              [  7];   //�ż�ȣ��             
	char offerjan                         [  6];   //�ŵ��ܷ�             
	char bidjan                           [  6];   //�ż��ܷ�             
	char S2offer                          [  7];   //���ŵ�ȣ��           
	char S2bid                            [  7];   //���ż�ȣ��           
	char S2offerjan                       [  6];   //���ŵ��ܷ�           
	char S2bidjan                         [  6];   //���ż��ܷ�           
	char S3offer                          [  7];   //�����ŵ�ȣ��         
	char S3bid                            [  7];   //�����ż�ȣ��         
	char S3offerjan                       [  6];   //�����ŵ��ܷ�         
	char S3bidjan                         [  6];   //�����ż��ܷ�         
	char S4offer                          [  7];   //4���ŵ�ȣ��          
	char S4bid                            [  7];   //4���ż�ȣ��          
	char S4offerjan                       [  6];   //4���ŵ��ܷ�          
	char S4bidjan                         [  6];   //4���ż��ܷ�          
	char S5offer                          [  7];   //5���ŵ�ȣ��          
	char S5bid                            [  7];   //5���ż�ȣ��          
	char S5offerjan                       [  6];   //5���ŵ��ܷ�          
	char S5bidjan                         [  6];   //5���ż��ܷ�          
	char S6offer                          [  7];   //6���ŵ�ȣ��          
	char S6bid                            [  7];   //6���ż�ȣ��          
	char S6offerjan                       [  6];   //6���ŵ��ܷ�          
	char S6bidjan                         [  6];   //6���ż��ܷ�          
	char S7offer                          [  7];   //7���ŵ�ȣ��          
	char S7bid                            [  7];   //7���ż�ȣ��          
	char S7offerjan                       [  6];   //7���ŵ��ܷ�          
	char S7bidjan                         [  6];   //7���ż��ܷ�          
	char S8offer                          [  7];   //8���ŵ�ȣ��          
	char S8bid                            [  7];   //8���ż�ȣ��          
	char S8offerjan                       [  6];   //8���ŵ��ܷ�          
	char S8bidjan                         [  6];   //8���ż��ܷ�          
	char S9offer                          [  7];   //9���ŵ�ȣ��          
	char S9bid                            [  7];   //9���ż�ȣ��          
	char S9offerjan                       [  6];   //9���ŵ��ܷ�          
	char S9bidjan                         [  6];   //9���ż��ܷ�          
	char S0offer                          [  7];   //10���ŵ�ȣ��         
	char S0bid                            [  7];   //10���ż�ȣ��         
	char S0offerjan                       [  6];   //10���ŵ��ܷ�         
	char S0bidjan                         [  6];   //10���ż��ܷ�         
	char offersu                          [  4];   //�� �� �ŵ� �Ǽ�              
	char bidsu                            [  4];   //�� �� �ż� �Ǽ�      
	char S2offersu                        [  4];   //�� �� �ŵ� �Ǽ�      
	char S2bidsu                          [  4];   //�� �� �ż� �Ǽ�      
	char S3offersu                        [  4];   //3���� �ŵ� �Ǽ�      
	char S3bidsu                          [  4];   //3���� �ż� �Ǽ�      
	char S4offersu                        [  4];   //4���� �ŵ� �Ǽ�      
	char S4bidsu                          [  4];   //4���� �ż� �Ǽ�      
	char S5offersu                        [  4];   //5���� �ŵ� �Ǽ�      
	char S5bidsu                          [  4];   //5���� �ż� �Ǽ�      
	char S6offersu                        [  4];   //6���� �ŵ� �Ǽ�      
	char S6bidsu                          [  4];   //6���� �ż� �Ǽ�      
	char S7offersu                        [  4];   //7���� �ŵ� �Ǽ�      
	char S7bidsu                          [  4];   //7���� �ż� �Ǽ�      
	char S8offersu                        [  4];   //8���� �ŵ� �Ǽ�      
	char S8bidsu                          [  4];   //8���� �ż� �Ǽ�      
	char S9offersu                        [  4];   //9���� �ŵ� �Ǽ�      
	char S9bidsu                          [  4];   //9���� �ż� �Ǽ�      
	char S0offersu                        [  4];   //10���� �ŵ� �Ǽ�     
	char S0bidsu                          [  4];   //10���� �ż� �Ǽ�     
	char tofferjan                        [  6];   //�Ѹŵ�ȣ�� �ܷ�      
	char tobidjan                         [  6];   //�Ѹż� ȣ�� �ܷ�     
	char toffersu                         [  5];   //�� �ŵ� �Ǽ�         
	char tbidsu                           [  5];   //�� �ż� �Ǽ�         
} TvHOutBlock;

typedef struct tagvH
{
	TvHInBlock                        vhinblock                             ;  //�Է� 
	TvHOutBlock                       vhoutblock                            ;  //��� 
} TvH;

typedef struct tagvCInBlock    //�Է�
{
	char fuitem                           [  8];	char _fuitem;                             //�����ڵ�             
} TvCInBlock;

typedef struct tagvCOutBlock    //���
{
	char fuitem                           [  8];	char _fuitem;                             //�����ڵ�             
	char futime                           [  8];	char _futime;                             //�ð� HH:MM:SS        
	char jgubun                           [  8];	char _jgubun;                             //����               
	char fusign                           [  1];	char _fusign;                             //���ϴ�� ��ȣ        
	char fuchange                         [  7];	char _fuchange;                           //���ϴ��             
	char fucurr                           [  7];	char _fucurr;                             //���簡               
	char fuopen                           [  7];	char _fuopen;                             //�ð�                 
	char fuhigh                           [  7];	char _fuhigh;                             //��                 
	char fulow                            [  7];	char _fulow;                              //����                 
	char fuvol                            [  6];	char _fuvol;                              //ü�����             
	char fuvolall                         [  7];	char _fuvolall;                           //���� ü�����        
	char fuvalall                         [ 12];	char _fuvalall;                           //�����ŷ����         
	char openyak                          [  7];	char _openyak;                            //�̰�����������       
	char jandatecnt                       [  3];	char _jandatecnt;                         //�����ϼ�             
	char offer                            [  7];	char _offer;                              //�ŵ�ȣ��             
	char bid                              [  7];	char _bid;                                //�ż�ȣ��             
	char offerjan                         [  6];	char _offerjan;                           //�ŵ��ܷ�             
	char bidjan                           [  6];	char _bidjan;                             //�ż��ܷ�             
	char S2offer                          [  7];	char _S2offer;                            //���ŵ�ȣ��           
	char S2bid                            [  7];	char _S2bid;                              //���ż�ȣ��           
	char S2offerjan                       [  6];	char _S2offerjan;                         //���ŵ��ܷ�           
	char S2bidjan                         [  6];	char _S2bidjan;                           //���ż��ܷ�           
	char S3offer                          [  7];	char _S3offer;                            //�����ŵ�ȣ��         
	char S3bid                            [  7];	char _S3bid;                              //�����ż�ȣ��         
	char S3offerjan                       [  6];	char _S3offerjan;                         //�����ŵ��ܷ�         
	char S3bidjan                         [  6];	char _S3bidjan;                           //�����ż��ܷ�         
	char S4offer                          [  7];	char _S4offer;                            //4���ŵ�ȣ��          
	char S4bid                            [  7];	char _S4bid;                              //4���ż�ȣ��          
	char S4offerjan                       [  6];	char _S4offerjan;                         //4���ŵ��ܷ�          
	char S4bidjan                         [  6];	char _S4bidjan;                           //4���ż��ܷ�          
	char S5offer                          [  7];	char _S5offer;                            //5���ŵ�ȣ��          
	char S5bid                            [  7];	char _S5bid;                              //5���ż�ȣ��          
	char S5offerjan                       [  6];	char _S5offerjan;                         //5���ŵ��ܷ�          
	char S5bidjan                         [  6];	char _S5bidjan;                           //5���ż��ܷ�          
	char S6offer                          [  7];	char _S6offer;                            //6���ŵ�ȣ��          
	char S6bid                            [  7];	char _S6bid;                              //6���ż�ȣ��          
	char S6offerjan                       [  6];	char _S6offerjan;                         //6���ŵ��ܷ�          
	char S6bidjan                         [  6];	char _S6bidjan;                           //6���ż��ܷ�          
	char S7offer                          [  7];	char _S7offer;                            //7���ŵ�ȣ��          
	char S7bid                            [  7];	char _S7bid;                              //7���ż�ȣ��          
	char S7offerjan                       [  6];	char _S7offerjan;                         //7���ŵ��ܷ�          
	char S7bidjan                         [  6];	char _S7bidjan;                           //7���ż��ܷ�          
	char S8offer                          [  7];	char _S8offer;                            //8���ŵ�ȣ��          
	char S8bid                            [  7];	char _S8bid;                              //8���ż�ȣ��          
	char S8offerjan                       [  6];	char _S8offerjan;                         //8���ŵ��ܷ�          
	char S8bidjan                         [  6];	char _S8bidjan;                           //8���ż��ܷ�          
	char S9offer                          [  7];	char _S9offer;                            //9���ŵ�ȣ��          
	char S9bid                            [  7];	char _S9bid;                              //9���ż�ȣ��          
	char S9offerjan                       [  6];	char _S9offerjan;                         //9���ŵ��ܷ�          
	char S9bidjan                         [  6];	char _S9bidjan;                           //9���ż��ܷ�          
	char S0offer                          [  7];	char _S0offer;                            //10���ŵ�ȣ��         
	char S0bid                            [  7];	char _S0bid;                              //10���ż�ȣ��         
	char S0offerjan                       [  6];	char _S0offerjan;                         //10���ŵ��ܷ�         
	char S0bidjan                         [  6];	char _S0bidjan;                           //10���ż��ܷ�         
	char offersu                          [  4];	char _offersu;                            //�� �� �ŵ� �Ǽ�      
	char bidsu                            [  4];	char _bidsu;                              //�� �� �ż� �Ǽ�      
	char S2offersu                        [  4];	char _S2offersu;                          //�� �� �ŵ� �Ǽ�      
	char S2bidsu                          [  4];	char _S2bidsu;                            //�� �� �ż� �Ǽ�      
	char S3offersu                        [  4];	char _S3offersu;                          //3���� �ŵ� �Ǽ�      
	char S3bidsu                          [  4];	char _S3bidsu;                            //3���� �ż� �Ǽ�      
	char S4offersu                        [  4];	char _S4offersu;                          //4���� �ŵ� �Ǽ�      
	char S4bidsu                          [  4];	char _S4bidsu;                            //4���� �ż� �Ǽ�      
	char S5offersu                        [  4];	char _S5offersu;                          //5���� �ŵ� �Ǽ�      
	char S5bidsu                          [  4];	char _S5bidsu;                            //5���� �ż� �Ǽ�      
	char S6offersu                        [  4];	char _S6offersu;                          //6���� �ŵ� �Ǽ�      
	char S6bidsu                          [  4];	char _S6bidsu;                            //6���� �ż� �Ǽ�      
	char S7offersu                        [  4];	char _S7offersu;                          //7���� �ŵ� �Ǽ�      
	char S7bidsu                          [  4];	char _S7bidsu;                            //7���� �ż� �Ǽ�      
	char S8offersu                        [  4];	char _S8offersu;                          //8���� �ŵ� �Ǽ�      
	char S8bidsu                          [  4];	char _S8bidsu;                            //8���� �ż� �Ǽ�      
	char S9offersu                        [  4];	char _S9offersu;                          //9���� �ŵ� �Ǽ�      
	char S9bidsu                          [  4];	char _S9bidsu;                            //9���� �ż� �Ǽ�      
	char S0offersu                        [  4];	char _S0offersu;                          //10���� �ŵ� �Ǽ�     
	char S0bidsu                          [  4];	char _S0bidsu;                            //10���� �ż� �Ǽ�     
	char tofferjan                        [  6];	char _tofferjan;                          //�Ѹŵ�ȣ�� �ܷ�      
	char tobidjan                         [  6];	char _tobidjan;                           //�Ѹż� ȣ�� �ܷ�     
	char toffersu                         [  5];	char _toffersu;                           //�� �ŵ� �Ǽ�         
	char tbidsu                           [  5];	char _tbidsu;                             //�� �ż� �Ǽ�         
	char chrate                           [  5];	char _chrate;                             //�����               
	char basis                            [  7];	char _basis;                              //���̽ý�             
	char grate                            [  7];	char _grate;                              //������               
	char gratio                           [  6];	char _gratio;                             //������               
	char preopenyak                       [  7];	char _preopenyak;                         //�̰����������Ϻ�     
	char bp_jgubun                        [  1];	char _bp_jgubun;                          //BP�� �屸��          
	char fspvolall                        [  7];	char _fspvolall;                          //�������� ü�����    
} TvCOutBlock;

typedef struct tagvC
{
	TvCInBlock                        vcinblock                             ;  //�Է� 
	TvCOutBlock                       vcoutblock                            ;  //��� 
} TvC;

typedef struct tagvVInBlock    //�Է�
{
	char fuitem                           [  8];   //�����ڵ�             
} TvVInBlock;

typedef struct tagvVOutBlock    //���
{
	char fuitem                           [  8];   //�����ڵ�             
	char theoryprice                      [  7];   //���� �̷а�          
	char theorytime                       [  8];   //�̷а� �ð�          
	char grate                            [  7];   //������               
	char gratio                           [  6];   //������               
} TvVOutBlock;

typedef struct tagvV
{
	TvVInBlock                        vvinblock                             ;  //�Է� 
	TvVOutBlock                       vvoutblock                            ;  //��� 
} TvV;

typedef struct tagvMInBlock    //�Է�
{
	char fuitem                           [  8];   //�����ڵ�             
} TvMInBlock;

typedef struct tagvMOutBlock    //���
{
	char fuitem                           [  8];   //�����ڵ�             
	char chetime                          [  8];   //ü��ð�             
	char openyak                          [  7];   //�̰�����������       
	char preopenyak                       [  7];   //���Ϲ̰�����������   
} TvMOutBlock;

typedef struct tagvM
{
	TvMInBlock                        vminblock                             ;  //�Է� 
	TvMOutBlock                       vmoutblock                            ;  //��� 
} TvM;

typedef struct tagv7InBlock    //�Է�
{
	char fspitem                          [  8];   //�������������ڵ�     
} Tv7InBlock;

typedef struct tagv7OutBlock    //���
{
	char fspitem                          [  8];   //�������������ڵ�     
	char fsptime                          [  8];   //�ð� HH:MM:SS        
	char offer                            [  8];   //�ŵ�ȣ��             
	char bid                              [  8];   //�ż�ȣ��             
	char offerjan                         [  6];   //�ŵ��ܷ�             
	char bidjan                           [  6];   //�ż��ܷ�             
	char S2offer                          [  8];   //���ŵ�ȣ��           
	char S2bid                            [  8];   //���ż�ȣ��           
	char S2offerjan                       [  6];   //���ŵ��ܷ�           
	char S2bidjan                         [  6];   //���ż��ܷ�           
	char S3offer                          [  8];   //�����ŵ�ȣ��         
	char S3bid                            [  8];   //�����ż�ȣ��         
	char S3offerjan                       [  6];   //�����ŵ��ܷ�         
	char S3bidjan                         [  6];   //�����ż��ܷ�         
	char S4offer                          [  8];   //4���ŵ�ȣ��          
	char S4bid                            [  8];   //4���ż�ȣ��          
	char S4offerjan                       [  6];   //4���ŵ��ܷ�          
	char S4bidjan                         [  6];   //4���ż��ܷ�          
	char S5offer                          [  8];   //5���ŵ�ȣ��          
	char S5bid                            [  8];   //5���ż�ȣ��          
	char S5offerjan                       [  6];   //5���ŵ��ܷ�          
	char S5bidjan                         [  6];   //5���ż��ܷ�          
	char S6offer                          [  8];   //6���ŵ�ȣ��          
	char S6bid                            [  8];   //6���ż�ȣ��          
	char S6offerjan                       [  6];   //6���ŵ��ܷ�          
	char S6bidjan                         [  6];   //6���ż��ܷ�          
	char S7offer                          [  8];   //7���ŵ�ȣ��          
	char S7bid                            [  8];   //7���ż�ȣ��          
	char S7offerjan                       [  6];   //7���ŵ��ܷ�          
	char S7bidjan                         [  6];   //7���ż��ܷ�          
	char S8offer                          [  8];   //8���ŵ�ȣ��          
	char S8bid                            [  8];   //8���ż�ȣ��          
	char S8offerjan                       [  6];   //8���ŵ��ܷ�          
	char S8bidjan                         [  6];   //8���ż��ܷ�          
	char S9offer                          [  8];   //9���ŵ�ȣ��          
	char S9bid                            [  8];   //9���ż�ȣ��          
	char S9offerjan                       [  6];   //9���ŵ��ܷ�          
	char S9bidjan                         [  6];   //9���ż��ܷ�          
	char S0offer                          [  8];   //10���ŵ�ȣ��         
	char S0bid                            [  8];   //10���ż�ȣ��         
	char S0offerjan                       [  6];   //10���ŵ��ܷ�         
	char S0bidjan                         [  6];   //10���ż��ܷ�         
	char offersu                          [  4];   //�� �� �ŵ� �Ǽ�      
	char bidsu                            [  4];   //�� �� �ż� �Ǽ�      
	char S2offersu                        [  4];   //�� �� �ŵ� �Ǽ�      
	char S2bidsu                          [  4];   //�� �� �ż� �Ǽ�      
	char S3offersu                        [  4];   //3���� �ŵ� �Ǽ�      
	char S3bidsu                          [  4];   //3���� �ż� �Ǽ�      
	char S4offersu                        [  4];   //4���� �ŵ� �Ǽ�      
	char S4bidsu                          [  4];   //4���� �ż� �Ǽ�      
	char S5offersu                        [  4];   //5���� �ŵ� �Ǽ�      
	char S5bidsu                          [  4];   //5���� �ż� �Ǽ�      
	char S6offersu                        [  4];   //6���� �ŵ� �Ǽ�      
	char S6bidsu                          [  4];   //6���� �ż� �Ǽ�      
	char S7offersu                        [  4];   //7���� �ŵ� �Ǽ�      
	char S7bidsu                          [  4];   //7���� �ż� �Ǽ�      
	char S8offersu                        [  4];   //8���� �ŵ� �Ǽ�      
	char S8bidsu                          [  4];   //8���� �ż� �Ǽ�      
	char S9offersu                        [  4];   //9���� �ŵ� �Ǽ�      
	char S9bidsu                          [  4];   //9���� �ż� �Ǽ�      
	char S0offersu                        [  4];   //10���� �ŵ� �Ǽ�     
	char S0bidsu                          [  4];   //10���� �ż� �Ǽ�     
	char tofferjan                        [  6];   //�Ѹŵ�ȣ�� �ܷ�      
	char tobidjan                         [  6];   //�Ѹż� ȣ�� �ܷ�     
	char toffersu                         [  5];   //�� �ŵ� �Ǽ�         
	char tbidsu                           [  5];   //�� �ż� �Ǽ�         
} Tv7OutBlock;

typedef struct tagv7
{
	Tv7InBlock                        v7inblock                             ;  //�Է� 
	Tv7OutBlock                       v7outblock                            ;  //��� 
} Tv7;

typedef struct tagv8InBlock    //�Է�
{
	char fspitem                          [  8];	char _fspitem;                            //�������������ڵ�     
} Tv8InBlock;

typedef struct tagv8OutBlock    //���
{
	char fspitem                          [  8];	char _fspitem;                            //�������������ڵ�     
	char fsptime                          [  8];	char _fsptime;                            //�ð� HH:MM:SS        
	char jgubun                           [  8];	char _jgubun;                             //����               
	char fspsign                          [  1];	char _fspsign;                            //���ϴ�� ��ȣ        
	char fspchange                        [  7];	char _fspchange;                          //���ϴ��             
	char fspcurr                          [  8];	char _fspcurr;                            //���簡               
	char fspcurr1                         [  7];	char _fspcurr1;                           //����������-�ٿ���    
	char fspcurr2                         [  7];	char _fspcurr2;                           //����������-������    
	char fspopen                          [  8];	char _fspopen;                            //�ð�                 
	char fsphigh                          [  8];	char _fsphigh;                            //��                 
	char fsplow                           [  8];	char _fsplow;                             //����                 
	char fspvol                           [  6];	char _fspvol;                             //ü�����             
	char fspvolall                        [  7];	char _fspvolall;                          //���� ü�����        
	char fspvalall                        [ 12];	char _fspvalall;                          //�����ŷ����         
	char offer                            [  8];	char _offer;                              //�ŵ�ȣ��             
	char bid                              [  8];	char _bid;                                //�ż�ȣ��             
	char offerjan                         [  6];	char _offerjan;                           //�ŵ��ܷ�             
	char bidjan                           [  6];	char _bidjan;                             //�ż��ܷ�             
	char S2offer                          [  8];	char _S2offer;                            //���ŵ�ȣ��           
	char S2bid                            [  8];	char _S2bid;                              //���ż�ȣ��           
	char S2offerjan                       [  6];	char _S2offerjan;                         //���ŵ��ܷ�           
	char S2bidjan                         [  6];	char _S2bidjan;                           //���ż��ܷ�           
	char S3offer                          [  8];	char _S3offer;                            //�����ŵ�ȣ��         
	char S3bid                            [  8];	char _S3bid;                              //�����ż�ȣ��         
	char S3offerjan                       [  6];	char _S3offerjan;                         //�����ŵ��ܷ�         
	char S3bidjan                         [  6];	char _S3bidjan;                           //�����ż��ܷ�         
	char S4offer                          [  8];	char _S4offer;                            //4���ŵ�ȣ��          
	char S4bid                            [  8];	char _S4bid;                              //4���ż�ȣ��          
	char S4offerjan                       [  6];	char _S4offerjan;                         //4���ŵ��ܷ�          
	char S4bidjan                         [  6];	char _S4bidjan;                           //4���ż��ܷ�          
	char S5offer                          [  8];	char _S5offer;                            //5���ŵ�ȣ��          
	char S5bid                            [  8];	char _S5bid;                              //5���ż�ȣ��          
	char S5offerjan                       [  6];	char _S5offerjan;                         //5���ŵ��ܷ�          
	char S5bidjan                         [  6];	char _S5bidjan;                           //5���ż��ܷ�          
	char S6offer                          [  8];	char _S6offer;                            //6���ŵ�ȣ��          
	char S6bid                            [  8];	char _S6bid;                              //6���ż�ȣ��          
	char S6offerjan                       [  6];	char _S6offerjan;                         //6���ŵ��ܷ�          
	char S6bidjan                         [  6];	char _S6bidjan;                           //6���ż��ܷ�          
	char S7offer                          [  8];	char _S7offer;                            //7���ŵ�ȣ��          
	char S7bid                            [  8];	char _S7bid;                              //7���ż�ȣ��          
	char S7offerjan                       [  6];	char _S7offerjan;                         //7���ŵ��ܷ�          
	char S7bidjan                         [  6];	char _S7bidjan;                           //7���ż��ܷ�          
	char S8offer                          [  8];	char _S8offer;                            //8���ŵ�ȣ��          
	char S8bid                            [  8];	char _S8bid;                              //8���ż�ȣ��          
	char S8offerjan                       [  6];	char _S8offerjan;                         //8���ŵ��ܷ�          
	char S8bidjan                         [  6];	char _S8bidjan;                           //8���ż��ܷ�          
	char S9offer                          [  8];	char _S9offer;                            //9���ŵ�ȣ��          
	char S9bid                            [  8];	char _S9bid;                              //9���ż�ȣ��          
	char S9offerjan                       [  6];	char _S9offerjan;                         //9���ŵ��ܷ�          
	char S9bidjan                         [  6];	char _S9bidjan;                           //9���ż��ܷ�          
	char S0offer                          [  8];	char _S0offer;                            //10���ŵ�ȣ��         
	char S0bid                            [  8];	char _S0bid;                              //10���ż�ȣ��         
	char S0offerjan                       [  6];	char _S0offerjan;                         //10���ŵ��ܷ�         
	char S0bidjan                         [  6];	char _S0bidjan;                           //10���ż��ܷ�         
	char offersu                          [  4];	char _offersu;                            //�� �� �ŵ� �Ǽ�      
	char bidsu                            [  4];	char _bidsu;                              //�� �� �ż� �Ǽ�      
	char S2offersu                        [  4];	char _S2offersu;                          //�� �� �ŵ� �Ǽ�      
	char S2bidsu                          [  4];	char _S2bidsu;                            //�� �� �ż� �Ǽ�      
	char S3offersu                        [  4];	char _S3offersu;                          //3���� �ŵ� �Ǽ�      
	char S3bidsu                          [  4];	char _S3bidsu;                            //3���� �ż� �Ǽ�      
	char S4offersu                        [  4];	char _S4offersu;                          //4���� �ŵ� �Ǽ�      
	char S4bidsu                          [  4];	char _S4bidsu;                            //4���� �ż� �Ǽ�      
	char S5offersu                        [  4];	char _S5offersu;                          //5���� �ŵ� �Ǽ�      
	char S5bidsu                          [  4];	char _S5bidsu;                            //5���� �ż� �Ǽ�      
	char S6offersu                        [  4];	char _S6offersu;                          //6���� �ŵ� �Ǽ�      
	char S6bidsu                          [  4];	char _S6bidsu;                            //6���� �ż� �Ǽ�      
	char S7offersu                        [  4];	char _S7offersu;                          //7���� �ŵ� �Ǽ�      
	char S7bidsu                          [  4];	char _S7bidsu;                            //7���� �ż� �Ǽ�      
	char S8offersu                        [  4];	char _S8offersu;                          //8���� �ŵ� �Ǽ�      
	char S8bidsu                          [  4];	char _S8bidsu;                            //8���� �ż� �Ǽ�      
	char S9offersu                        [  4];	char _S9offersu;                          //9���� �ŵ� �Ǽ�      
	char S9bidsu                          [  4];	char _S9bidsu;                            //9���� �ż� �Ǽ�      
	char S0offersu                        [  4];	char _S0offersu;                          //10���� �ŵ� �Ǽ�     
	char S0bidsu                          [  4];	char _S0bidsu;                            //10���� �ż� �Ǽ�     
	char tofferjan                        [  6];	char _tofferjan;                          //�Ѹŵ�ȣ�� �ܷ�      
	char tobidjan                         [  6];	char _tobidjan;                           //�Ѹż� ȣ�� �ܷ�     
	char toffersu                         [  5];	char _toffersu;                           //�� �ŵ� �Ǽ�         
	char tbidsu                           [  5];	char _tbidsu;                             //�� �ż� �Ǽ�         
	char chrate                           [  5];	char _chrate;                             //�����               
	char bp_jgubun                        [  1];	char _bp_jgubun;                          //BP�� �屸��          
} Tv8OutBlock;

typedef struct tagv8
{
	Tv8InBlock                        v8inblock                             ;  //�Է� 
	Tv8OutBlock                       v8outblock                            ;  //��� 
} Tv8;

typedef struct tageCInBlock    //�Է�
{
	char code                             [  6];	char _code;                               //�����ڵ�             
} TeCInBlock;

typedef struct tageCOutBlock    //���
{
	char code                             [  6];	char _code;                               //���������ڵ�         
	char time                             [  8];	char _time;                               //�ð�(HH:MM:SS)       
	char price                            [  7];	char _price;                              //���簡               
	char sign                             [  1];	char _sign;                               //�����ȣ             
	char change                           [  6];	char _change;                             //�����               
	char chrate                           [  5];	char _chrate;                             //�����               
	char open                             [  7];	char _open;                               //�ð�                 
	char high                             [  7];	char _high;                               //��                 
	char low                              [  7];	char _low;                                //����                 
	char offer                            [  7];	char _offer;                              //�ŵ�ȣ��             
	char bid                              [  7];	char _bid;                                //�ż�ȣ��             
	char volume                           [  9];	char _volume;                             //�����ŷ���           
	char volrate                          [  6];	char _volrate;                            //�ŷ��� ���Ϻ�        
	char movolume                         [  8];	char _movolume;                           //�����ŷ���           
	char value                            [  9];	char _value;                              //�ŷ���� �鸸��      
	char janggubun                        [  1];	char _janggubun;                          //�屸��               
	char cbgubun                          [  1];	char _cbgubun;                            //CB����               
	char stop                             [  1];	char _stop;                               //STOP                 
	char grate                            [  6];	char _grate;                              //������ 9(6)          
	char gratio                           [  8];	char _gratio;                             //������S9(5)V9(2)     
	char lphold                           [  9];	char _lphold;                             //LP��������           
	char lprate                           [  5];	char _lprate;                             //LP������             
} TeCOutBlock;

typedef struct tageC
{
	TeCInBlock                        ecinblock                             ;  //�Է� 
	TeCOutBlock                       ecoutblock                            ;  //��� 
} TeC;

typedef struct tageHInBlock    //�Է�
{
	char code                             [  6];   //�����ڵ�             
} TeHInBlock;

typedef struct tageHOutBlock    //���
{
	char code                             [  6];   //�����ڵ�             
	char time                             [  8];   //�ð�                 
	char S1_off                           [  7];   //�ŵ�ȣ��             
	char S1_bid                           [  7];   //�ż�ȣ��             
	char S1_offrem                        [  9];   //�ŵ�ȣ�� �ܷ�        
	char S1_bidrem                        [  9];   //�ż�ȣ�� �ܷ�        
	char S2_off                           [  7];   //2���ŵ�ȣ��          
	char S2_bid                           [  7];   //2���ż�ȣ��          
	char S2_offrem                        [  9];   //2���ŵ�ȣ�� �ܷ�     
	char S2_bidrem                        [  9];   //2���ż�ȣ�� �ܷ�     
	char S3_off                           [  7];   //3���ŵ�ȣ��          
	char S3_bid                           [  7];   //3���ż�ȣ��          
	char S3_offrem                        [  9];   //3���ŵ�ȣ�� �ܷ�     
	char S3_bidrem                        [  9];   //3���ż�ȣ�� �ܷ�     
	char S4_off                           [  7];   //4���ŵ�ȣ��          
	char S4_bid                           [  7];   //4���ż�ȣ��          
	char S4_offrem                        [  9];   //4���ŵ�ȣ�� �ܷ�     
	char S4_bidrem                        [  9];   //4���ż�ȣ�� �ܷ�     
	char S5_off                           [  7];   //5���ŵ�ȣ��          
	char S5_bid                           [  7];   //5���ż�ȣ��          
	char S5_offrem                        [  9];   //5���ŵ�ȣ�� �ܷ�     
	char S5_bidrem                        [  9];   //5���ż�ȣ�� �ܷ�     
	char S6_off                           [  7];   //6���ŵ�ȣ��          
	char S6_bid                           [  7];   //6���ż�ȣ��          
	char S6_offrem                        [  9];   //6���ŵ�ȣ�� �ܷ�     
	char S6_bidrem                        [  9];   //6���ż�ȣ�� �ܷ�     
	char S7_off                           [  7];   //7���ŵ�ȣ��          
	char S7_bid                           [  7];   //7���ż�ȣ��          
	char S7_offrem                        [  9];   //7���ŵ�ȣ�� �ܷ�     
	char S7_bidrem                        [  9];   //7���ż�ȣ�� �ܷ�     
	char S8_off                           [  7];   //8���ŵ�ȣ��          
	char S8_bid                           [  7];   //8���ż�ȣ��          
	char S8_offrem                        [  9];   //8���ŵ�ȣ�� �ܷ�     
	char S8_bidrem                        [  9];   //8���ż�ȣ�� �ܷ�     
	char S9_off                           [  7];   //9���ŵ�ȣ��          
	char S9_bid                           [  7];   //9���ż�ȣ��          
	char S9_offrem                        [  9];   //9���ŵ�ȣ�� �ܷ�     
	char S9_bidrem                        [  9];   //9���ż�ȣ�� �ܷ�     
	char S10_off                          [  7];   //10���ŵ�ȣ��         
	char S10_bid                          [  7];   //10���ż�ȣ��         
	char S10_offrem                       [  9];   //10���ŵ�ȣ�� �ܷ�    
	char S10_bidrem                       [  9];   //10���ż�ȣ�� �ܷ�    
	char T_offrem                         [  9];   //�Ѹŵ�ȣ�� �ܷ�      
	char T_bidrem                         [  9];   //�Ѹż�ȣ�� �ܷ�      
	char dongsi                           [  1];   //���ñ���             
	char eqprice                          [  7];   //����ȣ���ÿ���ü�ᰡ 
	char sign                             [  1];   //�����ȣ             
	char change                           [  6];   //�����               
	char chrate                           [  5];   //�����               
	char eqvol                            [  9];   //����ȣ���ÿ���ü����� 
	char S1_lpoffrem                      [  9];   //1��LP�ŵ�ȣ�� �ܷ�   
	char S1_lpbidrem                      [  9];   //1��LP�ż�ȣ�� �ܷ�   
	char S2_lpoffrem                      [  9];   //2��LP�ŵ�ȣ�� �ܷ�   
	char S2_lpbidrem                      [  9];   //2��LP�ż�ȣ�� �ܷ�   
	char S3_lpoffrem                      [  9];   //3��LP�ŵ�ȣ�� �ܷ�   
	char S3_lpbidrem                      [  9];   //3��LP�ż�ȣ�� �ܷ�   
	char S4_lpoffrem                      [  9];   //4��LP�ŵ�ȣ�� �ܷ�   
	char S4_lpbidrem                      [  9];   //4��LP�ż�ȣ�� �ܷ�   
	char S5_lpoffrem                      [  9];   //5��LP�ŵ�ȣ�� �ܷ�   
	char S5_lpbidrem                      [  9];   //5��LP�ż�ȣ�� �ܷ�   
	char S6_lpoffrem                      [  9];   //6��LP�ŵ�ȣ�� �ܷ�   
	char S6_lpbidrem                      [  9];   //6��LP�ż�ȣ�� �ܷ�   
	char S7_lpoffrem                      [  9];   //7��LP�ŵ�ȣ�� �ܷ�   
	char S7_lpbidrem                      [  9];   //7��LP�ż�ȣ�� �ܷ�   
	char S8_lpoffrem                      [  9];   //8��LP�ŵ�ȣ�� �ܷ�   
	char S8_lpbidrem                      [  9];   //8��LP�ż�ȣ�� �ܷ�   
	char S9_lpoffrem                      [  9];   //9��LP�ŵ�ȣ�� �ܷ�   
	char S9_lpbidrem                      [  9];   //9��LP�ż�ȣ�� �ܷ�   
	char S10_lpoffrem                     [  9];   //10��LP�ŵ�ȣ�� �ܷ�  
	char S10_lpbidrem                     [  9];   //10��LP�ż�ȣ�� �ܷ�  
} TeHOutBlock;

typedef struct tageH
{
	TeHInBlock                        ehinblock                             ;  //�Է� 
	TeHOutBlock                       ehoutblock                            ;  //��� 
} TeH;

typedef struct tageVInBlock    //�Է�
{
	char code                             [  6];   //�����ڵ�             
} TeVInBlock;

typedef struct tageVOutBlock    //���
{
	char code                             [  6];   //�����ڵ�             
	char theorytime                       [  8];   //�̷а��ð�           
	char theoryprice                      [  7];   //�̷а�               
	char impv                             [ 10];   //���纯����           
	char delta                            [  9];   //��ȣ+��Ÿ            
	char gmma                             [  9];   //��ȣ+����            
	char vega                             [  9];   //��ȣ+����            
	char theta                            [  9];   //��ȣ+��Ÿ            
	char rho                              [  9];   //��ȣ+��              
	char grate                            [  6];   //������               
	char gratio                           [  8];   //������               
} TeVOutBlock;

typedef struct tageV
{
	TeVInBlock                        evinblock                             ;  //�Է� 
	TeVOutBlock                       evoutblock                            ;  //��� 
} TeV;

typedef struct tageLInBlock    //�Է�
{
	char code                             [  6];   //�����ڵ�             
} TeLInBlock;

typedef struct tageLOutBlock    //���
{
	char code                             [  6];   //�����ڵ�             
	char jipyotime                        [  8];   //������ǥ�ð�         
	char parity                           [  8];   //�и�Ƽ               
	char egearing                         [  8];   //E��              
	char gearingrate                      [  8];   //������           
	char profitrate                       [  8];   //���ͺб���           
	char basepoint                        [  8];   //�ں�������           
	char filler                           [  6];   //FILLER               
} TeLOutBlock;

typedef struct tageL
{
	TeLInBlock                        elinblock                             ;  //�Է� 
	TeLOutBlock                       eloutblock                            ;  //��� 
} TeL;

typedef struct tageTInBlock    //�Է�
{
	char code                             [  6];	char _code;                               //�����ڵ�             
} TeTInBlock;

typedef struct tageTOutBlock    //���
{
	char code                             [  6];	char _code;                               //�����ڵ�             
	char time                             [  8];	char _time;                               //�ð�                 
	char off_trano1                       [  4];	char _off_trano1;                         //�ŵ�ȸ�����ڵ�1      
	char off_tra1                         [  6];	char _off_tra1;                           //�ŵ�ȸ������̸�1    
	char N_off_tra1                       [ 20];	char _N_off_tra1;                         //�ŵ�ȸ�������̸�1    
	char N_otraflag1                      [  1];	char _N_otraflag1;                        //�ŵ�ȸ����ܱ���1    
	char N_offvolume1                     [  9];	char _N_offvolume1;                       //�ŵ��ŷ���1          
	char N_offvolcha1                     [  9];	char _N_offvolcha1;                       //�����ŵ���1          	
	char bid_trano1                       [  4];	char _bid_trano1;                         //�ż�ȸ�����ڵ�1      
	char bid_tra1                         [  6];	char _bid_tra1;                           //�ż�ȸ������̸�1    
	char N_bid_tra1                       [ 20];	char _N_bid_tra1;                         //�ż�ȸ�������̸�1    
	char N_btraflag1                      [  1];	char _N_btraflag1;                        //�ż�ȸ����ܱ���1    
	char N_bidvolume1                     [  9];	char _N_bidvolume1;                       //�ż��ŷ���1          
	char N_bidvolcha1                     [  9];	char _N_bidvolcha1;                       //�����ż���1          
	char off_trano2                       [  4];	char _off_trano2;                         //�ŵ�ȸ�����ڵ�2      
	char off_tra2                         [  6];	char _off_tra2;                           //�ŵ�ȸ������̸�2    
	char N_off_tra2                       [ 20];	char _N_off_tra2;                         //�ŵ�ȸ�������̸�2    
	char N_otraflag2                      [  1];	char _N_otraflag2;                        //�ŵ�ȸ����ܱ���2    
	char N_offvolume2                     [  9];	char _N_offvolume2;                       //�ŵ��ŷ���2          
	char N_offvolcha2                     [  9];	char _N_offvolcha2;                       //�����ŵ���2          
	char bid_trano2                       [  4];	char _bid_trano2;                         //�ż�ȸ�����ڵ�2      
	char bid_tra2                         [  6];	char _bid_tra2;                           //�ż�ȸ������̸�2    
	char N_bid_tra2                       [ 20];	char _N_bid_tra2;                         //�ż�ȸ�������̸�2    
	char N_btraflag2                      [  1];	char _N_btraflag2;                        //�ż�ȸ����ܱ���2    
	char N_bidvolume2                     [  9];	char _N_bidvolume2;                       //�ż��ŷ���2          
	char N_bidvolcha2                     [  9];	char _N_bidvolcha2;                       //�����ż���2          
	char off_trano3                       [  4];	char _off_trano3;                         //�ŵ�ȸ�����ڵ�3      
	char off_tra3                         [  6];	char _off_tra3;                           //�ŵ�ȸ������̸�3    
	char N_off_tra3                       [ 20];	char _N_off_tra3;                         //�ŵ�ȸ�������̸�3    
	char N_otraflag3                      [  1];	char _N_otraflag3;                        //�ŵ�ȸ����ܱ���3    
	char N_offvolume3                     [  9];	char _N_offvolume3;                       //�ŵ��ŷ���3          
	char N_offvolcha3                     [  9];	char _N_offvolcha3;                       //�����ŵ���3          
	char bid_trano3                       [  4];	char _bid_trano3;                         //�ż�ȸ�����ڵ�3      
	char bid_tra3                         [  6];	char _bid_tra3;                           //�ż�ȸ������̸�3    
	char N_bid_tra3                       [ 20];	char _N_bid_tra3;                         //�ż�ȸ�������̸�3    
	char N_btraflag3                      [  1];	char _N_btraflag3;                        //�ż�ȸ����ܱ���3    
	char N_bidvolume3                     [  9];	char _N_bidvolume3;                       //�ż��ŷ���3          
	char N_bidvolcha3                     [  9];	char _N_bidvolcha3;                       //�����ż���3          
	char off_trano4                       [  4];	char _off_trano4;                         //�ŵ�ȸ�����ڵ�4      
	char off_tra4                         [  6];	char _off_tra4;                           //�ŵ�ȸ������̸�4    
	char N_off_tra4                       [ 20];	char _N_off_tra4;                         //�ŵ�ȸ�������̸�4    
	char N_otraflag4                      [  1];	char _N_otraflag4;                        //�ŵ�ȸ����ܱ���4    
	char N_offvolume4                     [  9];	char _N_offvolume4;                       //�ŵ��ŷ���4          
	char N_offvolcha4                     [  9];	char _N_offvolcha4;                       //�����ŵ���4          
	char bid_trano4                       [  4];	char _bid_trano4;                         //�ż�ȸ�����ڵ�4      
	char bid_tra4                         [  6];	char _bid_tra4;                           //�ż�ȸ������̸�4    
	char N_bid_tra4                       [ 20];	char _N_bid_tra4;                         //�ż�ȸ�������̸�4    
	char N_btraflag4                      [  1];	char _N_btraflag4;                        //�ż�ȸ����ܱ���4    
	char N_bidvolume4                     [  9];	char _N_bidvolume4;                       //�ż��ŷ���4          
	char N_bidvolcha4                     [  9];	char _N_bidvolcha4;                       //�����ż���4          
	char off_trano5                       [  4];	char _off_trano5;                         //�ŵ�ȸ�����ڵ�5      
	char off_tra5                         [  6];	char _off_tra5;                           //�ŵ�ȸ������̸�5    
	char N_off_tra5                       [ 20];	char _N_off_tra5;                         //�ŵ�ȸ�������̸�5    
	char N_otraflag5                      [  1];	char _N_otraflag5;                        //�ŵ�ȸ����ܱ���5    
	char N_offvolume5                     [  9];	char _N_offvolume5;                       //�ŵ��ŷ���5          
	char N_offvolcha5                     [  9];	char _N_offvolcha5;                       //�����ŵ���5          
	char bid_trano5                       [  4];	char _bid_trano5;                         //�ż�ȸ�����ڵ�5      
	char bid_tra5                         [  6];	char _bid_tra5;                           //�ż�ȸ������̸�5    
	char N_bid_tra5                       [ 20];	char _N_bid_tra5;                         //�ż�ȸ�������̸�5    
	char N_btraflag5                      [  1];	char _N_btraflag5;                        //�ż�ȸ����ܱ���5    
	char N_bidvolume5                     [  9];	char _N_bidvolume5;                       //�ż��ŷ���5          
	char N_bidvolcha5                     [  9];	char _N_bidvolcha5;                       //�����ż���5          
	char N_offvolall                      [  9];	char _N_offvolall;                        //�ܱ���ȸ����ŵ���   
	char N_offvolcha                      [  9];	char _N_offvolcha;                        //�ܱ��������ŵ���     
	char N_bidvolall                      [  9];	char _N_bidvolall;                        //�ܱ���ȸ����ż���   
	char N_bidvolcha                      [  9];	char _N_bidvolcha;                        //�ܱ��������ż���     
	char N_soonmaesu                      [  9];	char _N_soonmaesu;                        //�ܱ���ȸ�����ż�     
	char N_soonmaecha                     [  9];	char _N_soonmaecha;                       //�ܱ����������ż���   
	char N_alloffvol                      [  9];	char _N_alloffvol;                        //�ŵ���ü��           
	char N_allbidvol                      [  9];	char _N_allbidvol;                        //�ż���ü��           
	char hname                            [ 13];	char _hname;                              //�����               
	char kpgubun                          [  1];	char _kpgubun;                            //���屸��             
} TeTOutBlock;

typedef struct tageT
{
	TeTInBlock                        etinblock                             ;  //�Է� 
	TeTOutBlock                       etoutblock                            ;  //��� 
} TeT;

typedef struct tagfEInBlock    //�Է�
{
	char fuitem                           [  4];   //�����ڵ�             
} TfEInBlock;

typedef struct tagfEOutBlock    //���
{
	char fuitem                           [  4];   //�����ڵ�             
	char time                             [  8];   //�ð�                 
	char dongsi                           [  1];   //����ȣ������         
	char eqsign                           [  1];   //��������ȣ         
	char eqprice                          [  5];   //����ü�ᰡ           
	char eqchange                         [  5];   //��������           
	char eqchrate                         [  5];   //��������           
} TfEOutBlock;

typedef struct tagfE
{
	TfEInBlock                        feinblock                             ;  //�Է� 
	TfEOutBlock                       feoutblock                            ;  //��� 
} TfE;

typedef struct tagoEInBlock    //�Է�
{
	char opitem                           [  8];   //�����ڵ�             
} ToEInBlock;

typedef struct tagoEOutBlock    //���
{
	char opitem                           [  8];   //�����ڵ�             
	char time                             [  8];   //�ð�                 
	char dongsi                           [  1];   //����ȣ������         
	char eqsign                           [  1];   //��������ȣ         
	char eqprice                          [  5];   //����ü�ᰡ           
	char eqchange                         [  5];   //��������           
	char eqchrate                         [  5];   //��������           
} ToEOutBlock;

typedef struct tagoE
{
	ToEInBlock                        oeinblock                             ;  //�Է� 
	ToEOutBlock                       oeoutblock                            ;  //��� 
} ToE;

typedef struct tagvEInBlock    //�Է�
{
	char expcode                          [  8];   //�����ڵ�             
} TvEInBlock;

typedef struct tagvEOutBlock    //���
{
	char expcode                          [  8];   //�����ڵ�             
	char time                             [  8];   //�ð�                 
	char dongsi                           [  1];   //����ȣ������         
	char eqsign                           [  1];   //��������ȣ         
	char eqprice                          [  7];   //����ü�ᰡ           
	char eqchange                         [  7];   //��������           
	char eqchrate                         [  5];   //��������           
} TvEOutBlock;

typedef struct tagvE
{
	TvEInBlock                        veinblock                             ;  //�Է� 
	TvEOutBlock                       veoutblock                            ;  //��� 
} TvE;

typedef struct tagf7InBlock    //�Է�
{
	char fuitem                           [  4];	char _fuitem;                             //�����ڵ�             
} Tf7InBlock;

typedef struct tagf7OutBlock    //���
{
	char fuitem                           [  4];	char _fuitem;                             //�����ڵ�             
	char futime                           [  8];	char _futime;                             //�ð�                 
	char exlmtstep                        [  1];	char _exlmtstep;                          //����Ȯ�뿹���ܰ�     
	char exlmtgb                          [  1];	char _exlmtgb;                            //����Ȯ�뿹�� ����    
	char uplmtgb                          [  1];	char _uplmtgb;                            //��������Ȯ����Ѵܰ� 
	char dnlmtgb                          [  1];	char _dnlmtgb;                            //��������Ȯ�����Ѵܰ� 
	char uplmtprice                       [  5];	char _uplmtprice;                         //����� �ܰ� ���Ѱ�   
	char dnlmtprice                       [  5];	char _dnlmtprice;                         //����� �ܰ� ���Ѱ�   
} Tf7OutBlock;

typedef struct tagf7
{
	Tf7InBlock                        f7inblock                             ;  //�Է� 
	Tf7OutBlock                       f7outblock                            ;  //��� 
} Tf7;

typedef struct tago7InBlock    //�Է�
{
	char opitem                           [  8];	char _opitem;                             //�����ڵ�             
} To7InBlock;

typedef struct tago7OutBlock    //���
{
	char opitem                           [  8];	char _opitem;                             //�����ڵ�             
	char optime                           [  8];	char _optime;                             //�ð�                 
	char uplmtgb                          [  1];	char _uplmtgb;                            //��������Ȯ����Ѵܰ� 
	char dnlmtgb                          [  1];	char _dnlmtgb;                            //��������Ȯ�����Ѵܰ� 
	char uplmtprice                       [  5];	char _uplmtprice;                         //����� �ܰ� ���Ѱ�   
	char dnlmtprice                       [  5];	char _dnlmtprice;                         //����� �ܰ� ���Ѱ�   
} To7OutBlock;

typedef struct tago7
{
	To7InBlock                        o7inblock                             ;  //�Է� 
	To7OutBlock                       o7outblock                            ;  //��� 
} To7;

typedef struct tagvIInBlock    //�Է�
{
	char fuitem                           [  8];	char _fuitem;                             //�����ڵ�             
} TvIInBlock;

typedef struct tagvIOutBlock    //���
{
	char fuitem                           [  8];	char _fuitem;                             //�����ڵ�             
	char futime                           [  8];	char _futime;                             //�ð�                 
	char exlmtstep                        [  1];	char _exlmtstep;                          //����Ȯ�뿹���ܰ�     
	char exlmtgb                          [  1];	char _exlmtgb;                            //����Ȯ�뿹�� ����    
	char uplmtgb                          [  1];	char _uplmtgb;                            //��������Ȯ����Ѵܰ� 
	char dnlmtgb                          [  1];	char _dnlmtgb;                            //��������Ȯ�����Ѵܰ� 
	char uplmtprice                       [  7];	char _uplmtprice;                         //����� �ܰ� ���Ѱ�   
	char dnlmtprice                       [  7];	char _dnlmtprice;                         //����� �ܰ� ���Ѱ�   
} TvIOutBlock;

typedef struct tagvI
{
	TvIInBlock                        viinblock                             ;  //�Է� 
	TvIOutBlock                       vioutblock                            ;  //��� 
} TvI;

typedef struct tagu1InBlock    //�Է�
{
	char jisucode                         [  2];	char _jisucode;                           //�����ڵ�             
} Tu1InBlock;

typedef struct tagu1OutBlock    //���
{
	char jisucode                         [  2];	char _jisucode;                           //�����ڵ�             
	char jisutime                         [  8];	char _jisutime;                           //�ð�                 
	char jisu                             [  8];	char _jisu;                               //����                 
	char jisusign                         [  1];	char _jisusign;                           //�����ȣ             
	char jisuchange                       [  8];	char _jisuchange;                         //�����               
	char jisuvolume                       [  8];	char _jisuvolume;                         //�ŷ���               
	char jisuvalue                        [  8];	char _jisuvalue;                          //�ŷ����             
	char jisuopen                         [  8];	char _jisuopen;                           //�ð�����             
	char jisuhigh                         [  8];	char _jisuhigh;                           //������             
	char jisuhightime                     [  8];	char _jisuhightime;                       //���ð�             
	char jisulow                          [  8];	char _jisulow;                            //��������             
	char jisulowtime                      [  8];	char _jisulowtime;                        //�����ð�             
	char jisuchrate                       [  5];	char _jisuchrate;                         //���������           
	char jisubrkvol                       [  5];	char _jisubrkvol;                         //�ŷ�����             
} Tu1OutBlock;

typedef struct tagu1
{
	Tu1InBlock                        u1inblock                             ;  //�Է� 
	Tu1OutBlock                       u1outblock                            ;  //��� 
} Tu1;

typedef struct tagk1InBlock    //�Է�
{
	char jisukcode                        [  2];	char _jisukcode;                          //�����ڵ�             
} Tk1InBlock;

typedef struct tagk1OutBlock    //���
{
	char jisukcode                        [  2];	char _jisukcode;                          //�����ڵ�             
	char jisuktime                        [  8];	char _jisuktime;                          //�ð�                 
	char jisuk                            [  8];	char _jisuk;                              //����                 
	char jisuksign                        [  1];	char _jisuksign;                          //�����ȣ             
	char jisukchange                      [  8];	char _jisukchange;                        //�����               
	char jisukvolume                      [  8];	char _jisukvolume;                        //�ŷ���               
	char jisukvalue                       [  8];	char _jisukvalue;                         //�ŷ����             
	char jisukopen                        [  8];	char _jisukopen;                          //�ð�����             
	char jisukhigh                        [  8];	char _jisukhigh;                          //������             
	char jisukhightime                    [  8];	char _jisukhightime;                      //���ð�             
	char jisuklow                         [  8];	char _jisuklow;                           //��������             
	char jisuklowtime                     [  8];	char _jisuklowtime;                       //�����ð�             
	char jisukchrate                      [  5];	char _jisukchrate;                        //���������           
	char jisukbrkvol                      [  5];	char _jisukbrkvol;                        //�ŷ�����             
} Tk1OutBlock;

typedef struct tagk1
{
	Tk1InBlock                        k1inblock                             ;  //�Է� 
	Tk1OutBlock                       k1outblock                            ;  //��� 
} Tk1;

