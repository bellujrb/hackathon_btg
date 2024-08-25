import 'package:flutter/material.dart';
import 'package:flutter_modular/flutter_modular.dart';
import 'package:frontend/core/styles/text_style.dart';
import 'package:google_fonts/google_fonts.dart';
// ignore: depend_on_referenced_packages
import 'package:intl/intl.dart';

class TokenApprovalScreen extends StatefulWidget {
  const TokenApprovalScreen({super.key});

  @override
  // ignore: library_private_types_in_public_api
  _TokenApprovalScreenState createState() => _TokenApprovalScreenState();
}

class _TokenApprovalScreenState extends State<TokenApprovalScreen> {
  final _quantityController = TextEditingController();
  final _valueController = TextEditingController();
  final _expiryDateController = TextEditingController();
  final _formKey = GlobalKey<FormState>();

  _selectDate(BuildContext context) async {
    DateTime? selectedDate = await showDatePicker(
      context: context,
      initialDate: DateTime.now(),
      firstDate: DateTime(2020),
      lastDate: DateTime(2030),
    );

    if (selectedDate != null) {
      setState(() {
        _expiryDateController.text =
            DateFormat('dd/MM/yyyy').format(selectedDate);
      });
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: const Color(0xFFF5F5F5),
      appBar: AppBar(
        leading: IconButton(
          icon: const Icon(Icons.arrow_back),
          onPressed: () {
            Modular.to.navigate("dashboard-selected");
          },
        ),
        backgroundColor: Colors.transparent,
        elevation: 0,
        iconTheme: const IconThemeData(color: Colors.black),
      ),
      body: Center(
        child: Container(
          padding: const EdgeInsets.all(24),
          margin: const EdgeInsets.symmetric(horizontal: 16),
          decoration: BoxDecoration(
            color: Colors.white,
            borderRadius: BorderRadius.circular(16),
            boxShadow: [
              BoxShadow(
                color: Colors.black.withOpacity(0.1),
                blurRadius: 10,
                offset: const Offset(0, 5),
              ),
            ],
          ),
          width: 600, 
          child: Form(
            key: _formKey,
            child: Column(
              mainAxisSize: MainAxisSize.min,
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Text("Aprovação de Token", style: GoogleFonts.poppins()),
                const SizedBox(height: 20),
                Text("Configuração do Token", style: GoogleFonts.poppins()),
                const SizedBox(height: 20),
                _buildTextField(
                  label: 'Quantidade de Tokens a Emitir',
                  controller: _quantityController,
                  hint: 'Digite o valor total ou frações do recebível',
                ),
                const SizedBox(height: 20),
                _buildTextField(
                  label: 'Valor do token',
                  controller: _valueController,
                  hint: 'Digite o valor nominal de cada token',
                ),
                const SizedBox(height: 20),
                _buildTextField(
                  label: 'Vencimento do Token',
                  controller: _expiryDateController,
                  hint: 'Data de vencimento para pagamento dos tokens',
                  readOnly: true,
                  onTap: () => _selectDate(context),
                ),
                const SizedBox(height: 30),
                SizedBox(
                  width: double.infinity,
                  child: ElevatedButton(
                    onPressed: () {
                      if (_formKey.currentState!.validate()) {
                        Modular.to.navigate("token-2");
                      }
                    },
                    style: ElevatedButton.styleFrom(
                      padding: const EdgeInsets.symmetric(vertical: 16),
                      backgroundColor: Colors.blue,
                      shape: RoundedRectangleBorder(
                        borderRadius: BorderRadius.circular(8),
                      ),
                    ),
                    child: Text(
                      'Continuar',
                      style: GoogleFonts.poppins(
                          textStyle: context.appTextStyles.smallWhite),
                    ),
                  ),
                ),
              ],
            ),
          ),
        ),
      ),
    );
  }

  Widget _buildTextField({
    required String label,
    required TextEditingController controller,
    String? hint,
    bool readOnly = false,
    VoidCallback? onTap,
  }) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Text(
          label,
          style: GoogleFonts.poppins()),
        const SizedBox(height: 10),
        TextFormField(
          controller: controller,
          readOnly: readOnly,
          decoration: InputDecoration(
            hintText: hint,
            border: OutlineInputBorder(
              borderRadius: BorderRadius.circular(8),
            ),
          ),
          onTap: onTap,
          validator: (value) {
            if (value == null || value.isEmpty) {
              return 'Campo obrigatório';
            }
            return null;
          },
        ),
      ],
    );
  }
}
