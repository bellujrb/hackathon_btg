import 'package:flutter/material.dart';
import 'package:flutter_modular/flutter_modular.dart';
import 'package:frontend/core/styles/colors.dart';
import 'package:frontend/core/styles/text_style.dart';
import 'package:google_fonts/google_fonts.dart';

class HomePage extends StatefulWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  // ignore: library_private_types_in_public_api
  _HomePageState createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  String selectedMenu = 'Início'; 
  String hoveredMenu = ''; 

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        
      ),
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
                      const SizedBox(
                        height: 6,
                      ),
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
                      const SizedBox(
                        height: 6,
                      ),
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
          Padding(
            padding: const EdgeInsets.all(24.0),
            child: Column(
              mainAxisAlignment: MainAxisAlignment.start,
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Text("Lista de verificação de integração", style: GoogleFonts.poppins(
                  textStyle: context.appTextStyles.smallBlack
                )),
                const SizedBox(
                  height: 10,
                ),
                SizedBox(
                  height: 128,
                  width: 336,
                  child: Image.asset('assets/checklist.png'),
                ),
                const SizedBox(
                  height: 20,
                ),
                Text("Ativos Financeiros Pendentes", style: GoogleFonts.poppins(
                  textStyle: context.appTextStyles.smallBlack
                )),
                const SizedBox(
                  height: 10,
                ),
                InkWell(
                  onTap: () => Modular.to.navigate("dashboard-selected"),
                  child: SizedBox(
                    height: 240,
                    width: 336,
                    child: Image.asset('assets/btg.png'),
                  ),
                ),
              ],
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
          tileColor:
              (isSelected || isHovered) ? Colors.white : Colors.transparent,
          shape: RoundedRectangleBorder(
            borderRadius: BorderRadius.circular(10), 
          ),
          onTap: onTap,
        ),
      ),
    );
  }
}
