import 'package:flutter/material.dart';
import 'package:flutter_modular/flutter_modular.dart';
import 'package:frontend/core/styles/colors.dart';
import 'package:frontend/core/styles/text_style.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:syncfusion_flutter_charts/charts.dart';

class DashboardSelected extends StatefulWidget {
  const DashboardSelected({Key? key}) : super(key: key);

  @override
  // ignore: library_private_types_in_public_api
  _DashboardSelectedState createState() => _DashboardSelectedState();
}

class _DashboardSelectedState extends State<DashboardSelected> {
  String selectedMenu = 'Início'; 
  String hoveredMenu = ''; 

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(),
      body: Row(
        children: [
          Container(
            width: 250,
            color: Colors.white,
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                const SizedBox(height: 50),
                const Padding(
                  padding: EdgeInsets.all(16.0),
                  child: Text(
                    'LUMQI',
                    style: TextStyle(
                      fontSize: 24,
                      fontWeight: FontWeight.bold,
                    ),
                  ),
                ),
                const SizedBox(height: 20),
                Padding(
                  padding: const EdgeInsets.all(16.0),
                  child: Column(
                    children: [
                      buildMenuItem(
                        icon: Icons.home,
                        text: 'Início',
                        isSelected: selectedMenu == 'Início',
                        isHovered: hoveredMenu == 'Início',
                        onTap: () {
                          setState(() {
                            selectedMenu = 'Início';
                          });
                        },
                      ),
                      const SizedBox(height: 6),
                      buildMenuItem(
                        icon: Icons.analytics,
                        text: 'Análise',
                        isSelected: selectedMenu == 'Análise',
                        isHovered: hoveredMenu == 'Análise',
                        onTap: () {
                          setState(() {
                            selectedMenu = 'Análise';
                          });
                        },
                      ),
                      const SizedBox(height: 6),
                      buildMenuItem(
                        icon: Icons.help_outline,
                        text: 'Ajuda',
                        isSelected: selectedMenu == 'Ajuda',
                        isHovered: hoveredMenu == 'Ajuda',
                        onTap: () {
                          setState(() {
                            selectedMenu = 'Ajuda';
                          });
                        },
                      ),
                    ],
                  ),
                )
              ],
            ),
          ),
          Expanded(
            child: Padding(
              padding: const EdgeInsets.all(24.0),
              child: SingleChildScrollView(
                child: Column(
                  mainAxisAlignment: MainAxisAlignment.start,
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Text(
                      "Lista de verificação de integração",
                      style: GoogleFonts.poppins(
                          textStyle: context.appTextStyles.smallBlack),
                    ),
                    const SizedBox(
                      height: 10,
                    ),
                    SizedBox(
                      height: 200,
                      width: 336,
                      child: Image.asset('assets/checklist2.png'),
                    ),
                    const SizedBox(
                      height: 10,
                    ),
                    InkWell(
                      onTap: () => Modular.to.navigate("token-1"),
                      child: SizedBox(
                        height: 390,
                        width: MediaQuery.of(context).size.width * 0.6,
                        child: Image.asset('assets/btg2.png'),
                      ),
                    ),
                    const SizedBox(
                      height: 10,
                    ),
                    SizedBox(
                      height: 184,
                      width: MediaQuery.of(context).size.width * 0.6,
                      child: Image.asset('assets/actions.png'),
                    ),
                    const SizedBox(
                      height: 20,
                    ),
                    Text(
                      "Análise Preditiva",
                      style: GoogleFonts.poppins(
                          textStyle: context.appTextStyles.smallBlack),
                    ),
                    const SizedBox(height: 20),
                    Row(
                      children: [
                        Expanded(
                          child: Card(
                            elevation: 4,
                            shape: RoundedRectangleBorder(
                              borderRadius: BorderRadius.circular(12),
                            ),
                            child: Padding(
                              padding: const EdgeInsets.all(16.0),
                              child: Column(
                                children: [
                                  Text(
                                    "Probabilidade de Aprovação",
                                    style: GoogleFonts.poppins(
                                        textStyle: context
                                            .appTextStyles.smallBlack
                                            .copyWith(fontSize: 16)),
                                  ),
                                  const SizedBox(height: 10),
                                  SizedBox(
                                    height: 200,
                                    child: SfCircularChart(
                                      legend: const Legend(isVisible: true),
                                      series: <CircularSeries>[
                                        PieSeries<ChartData, String>(
                                          dataSource: _getPieData(),
                                          xValueMapper: (ChartData data, _) =>
                                              data.category,
                                          yValueMapper: (ChartData data, _) =>
                                              data.value,
                                          dataLabelSettings: const DataLabelSettings(
                                              isVisible: true),
                                        ),
                                      ],
                                    ),
                                  ),
                                ],
                              ),
                            ),
                          ),
                        ),
                        const SizedBox(width: 20),
                        Expanded(
                          child: Card(
                            elevation: 4,
                            shape: RoundedRectangleBorder(
                              borderRadius: BorderRadius.circular(12),
                            ),
                            child: Padding(
                              padding: const EdgeInsets.all(16.0),
                              child: Column(
                                children: [
                                  Text(
                                    "Indicador de Risco",style: GoogleFonts.poppins(
                                        textStyle: context
                                            .appTextStyles.smallBlack
                                            .copyWith(fontSize: 16)),
                                  ),
                                  const SizedBox(height: 10),
                                  SizedBox(
                                    height: 200,
                                    child: SfCartesianChart(
                                      primaryXAxis: CategoryAxis(),
                                      series: <CartesianSeries>[
                                        BarSeries<ChartData, String>(
                                          dataSource: _getBarData(),
                                          xValueMapper: (ChartData data, _) =>
                                              data.category,
                                          yValueMapper: (ChartData data, _) =>
                                              data.value,
                                          dataLabelSettings: const DataLabelSettings(
                                              isVisible: true),
                                        )
                                      ],
                                    ),
                                  ),
                                ],
                              ),
                            ),
                          ),
                        ),
                      ],
                    ),
                  ],
                ),
              ),
            ),
          )
        ],
      ),
    );
  }

  Widget buildMenuItem({
    required IconData icon,
    required String text,
    required bool isSelected,
    required bool isHovered,
    required Function() onTap,
  }) {
    return MouseRegion(
      onEnter: (_) {
        setState(() {
          hoveredMenu = text;
        });
      },
      onExit: (_) {
        setState(() {
          hoveredMenu = '';
        });
      },
      child: Container(
        decoration: BoxDecoration(
            color: (isSelected || isHovered)
                ? AppColors.primary
                : Colors.transparent,
            borderRadius: BorderRadius.circular(16)),
        padding: const EdgeInsets.symmetric(vertical: 8.0),
        child: ListTile(
          leading: Icon(
            icon,
            color: (isSelected || isHovered) ? Colors.white : Colors.black54,
          ),
          title: Text(
            text,
            style: TextStyle(
              color: (isSelected || isHovered) ? Colors.white : Colors.black54,
              fontWeight: isSelected ? FontWeight.bold : FontWeight.normal,
            ),
          ),
          onTap: onTap,
        ),
      ),
    );
  }

  List<ChartData> _getPieData() {
    final List<ChartData> chartData = [
      ChartData('Aprovado', 75),
      ChartData('Rejeitado', 25),
    ];
    return chartData;
  }

  List<ChartData> _getBarData() {
    final List<ChartData> chartData = [
      ChartData('Baixo', 8),
      ChartData('Médio', 4),
      ChartData('Alto', 2),
    ];
    return chartData;
  }
}

class ChartData {
  final String category;
  final double value;

  ChartData(this.category, this.value);
}
